from xml.etree import ElementTree
import os
import requests
import zipfile

namespaces = {"xsd": "http://www.w3.org/2001/XMLSchema"}
schema_file = "VDV453_incl_454_V2017e.xsd"
schema_file = "VDV453_incl_454_V3.1.0_v12.xsd"

if not os.path.exists(schema_file):
    url = "https://www.vdv.de"
    zip_file = "vdv453-incl-454-v3.1.0-v12.zipx"
    response = requests.get(f"{url}/{zip_file}", stream=True)
    with open(zip_file, "wb") as f:
        for chunk in response.iter_content(chunk_size=8192):
            f.write(chunk)
    with zipfile.ZipFile(zip_file, "r") as zip_ref:
        zip_ref.extract(schema_file)
    os.unlink(zip_file)

with open(schema_file) as f:
    schema = ElementTree.parse(f)

xsd2go = {
    "xsd:boolean": "bool",
    "xsd:positiveInteger": "int",
    "FoStringType": "string",
    "xsd:NMTOKEN": "string",
    "xsd:string": "string",
    "xsd:nonNegativeInteger": "int",
    "xsd:integer": "int",
    "xsd:dateTime": "time.Time",
    "xsd:date": "time.Time"
}

def _get_type(type: str):
    if type.endswith("Enum"):
        return "string"
    elif type.endswith("Type"):
        return type[:-4]
    else:
        return xsd2go.get(type, type)

def parse_sequence(el: ElementTree.Element):
    for child in el:
        if child.tag.endswith('element'):
            name = child.attrib['name']
            type = child.attrib.get('type')
            if type is None:
                continue
            type = _get_type(type)
            print(f'  {name} []{type} `xml:"{name}"`')
        else:
            pass
            # raise Exception(child.tag, el.attrib['name'])

def parse_complex_type(el: ElementTree.Element):
    print("type " + _get_type(el.attrib['name']) + " struct {")
    for child in el:
        if child.tag.endswith('attribute'):
            name = child.attrib['name']
            type = child.attrib['type']
            type = _get_type(type)
            print(f'  {name} {type} `xml:"{name},attr"`')
        elif child.tag.endswith('sequence'):
            parse_sequence(child)
        elif child.tag.endswith('choice'):
            if el.attrib['name'] == "FahrtRefType":
                parse_sequence(next(iter(child)))
            else:
                for i in child:
                    parse_sequence(i)
        elif child.tag.endswith('annotation'):
            pass
        elif child.tag.endswith('group'):
            pass
        elif child.tag.endswith('simpleContent'):
            pass
        elif child.tag.endswith('complexContent'):
            pass
        else:
            raise Exception(child.tag)
    print("}")

print("package main")
print('import "time"')

for child in schema.getroot():
    if child.tag.endswith("annotation") or child.tag.endswith("import") or child.tag.endswith("element"):
        continue
    name = child.attrib["name"]
    if child.tag.endswith("simpleType"):
        if name.endswith("Enum"):
            name = name[:-4]
            for i in child.findall(".//xsd:enumeration", namespaces):
                const_name = f"{name}{i.attrib['value']}"
                const_name = const_name.replace("-", "") # For FoFahrzeugAusstattungsCodeUSB-Ladebuchse
                print(f'const {const_name} = "{const_name}"')
        elif name.endswith("Type"):
            name = name[:-4]
            assert len(child) == 1
            assert child[0].tag.endswith("restriction")
            base_type = child[0].attrib["base"]
            print(f"type {name} = {_get_type(base_type)}")
        else:
            raise Exception(name)
    elif child.tag.endswith("complexType") or child.tag.endswith("group"):
        parse_complex_type(child)
        pass
    else:
        raise Exception(child.tag)