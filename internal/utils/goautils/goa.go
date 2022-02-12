package goautils

import "go/ast"

// https://pkg.go.dev/goa.design/goa@v1.4.3/design/apidsl

// CheckTargetFunctionList returns function list that can use
func CheckTargetFunctionList() []string {
	// List functions that have "function" or "interface" as arguments
	funcs := []string{
		"API",            // func API(name string, dsl func()) *design.APIDefinition
		"APIKeySecurity", // func APIKeySecurity(name string, dsl ...func()) *design.SecuritySchemeDefinition
		//"AccessCodeFlow",      // func AccessCodeFlow(authorizationURL, tokenURL string)
		"Action", // func Action(name string, dsl func())
		//"ApplicationFlow",     // func ApplicationFlow(tokenURL string)
		"ArrayOf",    // func ArrayOf(v interface{}, dsl ...func()) *design.Array
		"Attribute",  // func Attribute(name string, args ...interface{})
		"Attributes", // func Attributes(apidsl func())
		//"BasePath",            // func BasePath(val string)
		"BasicAuthSecurity", // func BasicAuthSecurity(name string, dsl ...func()) *design.SecuritySchemeDefinition
		"CONNECT",           // func CONNECT(path string, dsl ...func()) *design.RouteDefinition
		//"CanonicalActionName", // func CanonicalActionName(a string)
		"CollectionOf", // func CollectionOf(v interface{}, paramAndDSL ...interface{}) *design.MediaTypeDefinition
		"Consumes",     // func Consumes(args ...interface{})
		"Contact",      // func Contact(dsl func())
		//"ContentType",         // func ContentType(typ string)
		//"Credentials", // func Credentials()
		"DELETE",       // func DELETE(path string, dsl ...func()) *design.RouteDefinitio
		"Default",      // func Default(def interface{})
		"DefaultMedia", // func DefaultMedia(val interface{}, viewName ...string)
		//"Description", // func Description(d string)
		"Docs", // func Docs(dsl func())
		//"Email",   // func Email(email string)
		"Enum",    // func Enum(val ...interface{})
		"Example", // func Example(exp interface{})
		//"Expose",  // func Expose(vals ...string)
		"Files", // func Files(path, filename string, dsls ...func())
		//"Format",   // func Format(f string)
		//"Function", // func Function(fn string)
		"GET",     // func GET(path string, dsl ...func()) *design.RouteDefinition
		"HEAD",    // func HEAD(path string, dsl ...func()) *design.RouteDefinition
		"HashOf",  // func HashOf(k, v interface{}, dsls ...func()) *design.Hash
		"Header",  // func Header(name string, args ...interface{})
		"Headers", // func Headers(params ...interface{})
		//"Host",  // func Host(host string)
		//"ImplicitFlow" // func ImplicitFlow(authorizationURL string)
		"JWTSecurity", // func JWTSecurity(name string, dsl ...func()) *design.SecuritySchemeDefinition
		"License",     // func License(dsl func())
		//"Link",            // func Link(name string, view ...string)
		"Links", // func Links(apidsl func())
		//"MaxAge",    // func MaxAge(val uint)
		//"MaxLength", // func MaxLength(val int)
		"Maximum",   // func Maximum(val interface{})
		"Media",     // func Media(val interface{}, viewName ...string)
		"MediaType", // func MediaType(identifier string, apidsl func()) *design.MediaTypeDefinition
		"Member",    // func Member(name string, args ...interface{})
		//"Metadata",  // func Metadata(name string, value ...string)
		//"Methods",   // func Methods(vals ...string)
		//"MinLength", // func MinLength(val int)
		//"Minimum",   // func Minimum(val interface{})
		//"MultipartForm", // func MultipartForm()
		//"Name", // func Name(name string)
		//"NoExample",      // func NoExample()
		"OAuth2Security",  // func OAuth2Security(name string, dsl ...func()) *design.SecuritySchemeDefinition
		"OPTIONS",         // func OPTIONS(path string, dsl ...func()) *design.RouteDefinition
		"OptionalPayload", // func OptionalPayload(p interface{}, dsls ...func())
		"Origin",          // func Origin(origin string, dsl func())
		"PATCH",           // func PATCH(path string, dsl ...func()) *design.RouteDefinition
		"POST",            // func POST(path string, dsl ...func()) *design.RouteDefinition
		"PUT",             // func PUT(path string, dsl ...func()) *design.RouteDefinition
		//"Package",         // func Package(path string)
		"Param",  // func Param(name string, args ...interface{})
		"Params", // func Params(dsl func())
		//"Parent",       // func Parent(p string)
		//"PasswordFlow", // func PasswordFlow(tokenURL string)
		//"Pattern", // func Pattern(p string)
		"Payload",  // func Payload(p interface{}, dsls ...func())
		"Produces", // func Produces(args ...interface{})
		//"Query", // func Query(parameterName string)
		//"ReadOnly", // func ReadOnly()
		//"Reference", // func Reference(t design.DataType)
		"Required",         // func Required(names ...string)
		"Resource",         // func Resource(name string, dsl func()) *design.ResourceDefinition
		"Response",         // func Response(name string, paramsAndDSL ...interface{})
		"ResponseTemplate", // func ResponseTemplate(name string, p interface{})
		//"Routing",          // func Routing(routes ...*design.RouteDefinition)
		//"Scheme",           // func Scheme(vals ...string)
		//"Scope", // func Scope(name string, desc ...string)
		"Security", // func Security(scheme interface{}, dsl ...func())
		//"Status",   // func Status(status int)
		"TRACE", // func TRACE(path string, dsl ...func()) *design.RouteDefinition
		//"TermsOfService", // func TermsOfService(terms string)
		//"Title",// func Title(val string)
		//"TokenURL", // func TokenURL(tokenURL string)
		"Trait", // func Trait(name string, val ...func())
		"Type",  // func Type(name string, dsl func()) *design.UserTypeDefinition
		//"TypeName", // func TypeName(name string)
		//"URL", // func URL(url string)
		//"UseTrait", //func UseTrait(names ...string)
		//"Version", // func Version(ver string)
		"View", // func View(name string, apidsl ...func())
	}
	return funcs
}

// NotWarnSyntaxCheck returns true if no warning is required in the syntax check.
func NotWarnSyntaxCheck(args []ast.Expr, funcName string) bool {
	// Functions with variadic arguments
	funcs := []string{
		"APIKeySecurity", "ArrayOf", "Attribute", "BasicAuthSecurity", "CONNECT",
		"CollectionOf", "Consumes", "DELETE", "DefaultMedia", "Enum", "Expose",
		"Files", "GET", "HEAD", "HashOf", "Header", "Headers", "JWTSecurity",
		"Media", "MediaType", "Member", "Methods", "OAuth2Security", "OPTIONS",
		"OptionalPayload", "PATCH", "POST", "PUT", "Param", "Payload", "Produces",
		"Response", "Routing", "Security", "TRACE", "Trait", "UseTrait", "View",
	}

	if len(args) > 1 {
		return false
	}

	for _, f := range funcs {
		if funcName == f {
			return true
		}
	}
	return false
}
