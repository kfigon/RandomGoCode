# Architectures:
* monolith 
    * selfcontained, fast
    * big codebase, long code change-test iterations
    * no sharing, 1 technology, uneven resource distribution
* SOA - services + service bus to integrate. Unified protocol - SOAP + WSDL
    * enables sharing and various technologies
    * large, complex, expensive ESB (issues with collaboration)
    * lack of tooling
    * ESB is too smart
* microservices - smart services, dumb pipes