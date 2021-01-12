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

## microservices
* componentization with good API in opposition to libraries in a monolith
* organization around business capabilities - we devide microservices around business functions (order service, transaction service). We also have small teams, which do everything - UI,logic,db. In monoilith these can be different teams
* product not project - outcome is business product, not 'my part'. Increased customer satisfaction
* smart endpoints, dumb pipes - logic is present in decentralized services, not in ESB. Integration is simple (no wsdl). SImple protocol! More chaos is good! It enhances agile mindset
	* direct connection is not always good idea - gateway or discovery service might be better
	* sometimes graphQL or gRPC might be better than REST
* decentralized governance - more agile approach to technology and tech decision (theoretically - logging, db, protocol, tools). More optimal tech decisions
* decentralized data managements - multiple DBs - encourage isolation and we can choose our own DB. Controversial topic
	* Not always possible
	* distributed transactions?
	* potential data duplication
* infrastructure automation - CI/CD tools, container orchestration. Not relying on ESB. Short deployment cycle is a must, use lots of tools and automation
* design for failure - lot of network traffic, so many things can go wrong. 
	* Gracefull error handling
	* logging and monitoring are required
	* caching in case of error
* evolutionary design - moving from monolith to microservices should be slow and gradual. Not possible to rewrite everything

## probles solved by microservices
