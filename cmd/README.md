# Composing the services
This document explains the design dicesions made for grouping and composing the microservices from the software.

## FAQ
* > Q: Should the projector and query handlers for a read model be in the same service?
  >
  > A: No. The projector and query handlers will scale independently. The projector load is related to the write load and not the read load.