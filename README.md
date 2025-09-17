# VDV 454

This is a simple VDV 454 implementation in Go. It is used the exchange data among different partners in public transport systems.

The VDV spec can be found at https://www.vdv.de/454v2.0-sd.pdfx. It relies on the underlying VDV 453 spec, which can be found at https://www.vdv.de/453v2.6-sds.pdfx.

## Usage

The main entry point of this library is the `VdvServer` class. It implements a basic server which can both subscribe for data from other server and distribute data to other servers.

