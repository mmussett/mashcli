package plandesigner

import (

	"github.com/mmussett/mashcli/cli/app/endpoints"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/mmussett/mashcli/cli/app/methods"
	"github.com/mmussett/mashcli/cli/app/packages"
	"github.com/mmussett/mashcli/cli/app/plans"
	"github.com/mmussett/mashcli/cli/app/planservices"
	"github.com/mmussett/mashcli/cli/app/services"
	"github.com/pkg/errors"
)

func AddServiceToPackagePlan(accessToken, serviceId, packageId, planId string) error {

	// Does Package exist?
	targetPackage, err := packages.Get(accessToken, &packages.MethodParams{PackageId: packageId}, &mashcli.Params{Fields: "id,name"})
	if err != nil {
		return err
	}


	// Does Plan exist?
	targetPlan, err := plans.Get(accessToken, &plans.MethodParams{PackageId: packageId, PlanId: planId}, &mashcli.Params{Fields: "id,name"})
	if err != nil {
		return err
	}

	// Get Service we're going to add to the Package-Plan
	sourceService, err := services.Get(accessToken, &services.MethodParams{ServiceId: serviceId}, &mashcli.Params{Fields: services.SERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	var currentPlanServices *[]planservices.PlanServices
	currentPlanServices, err = planservices.GetCollection(accessToken, &planservices.MethodParams{PackageId: packageId, PlanId: planId}, &mashcli.Params{Fields: planservices.PLANSERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	// Check that service isn't already added (the PlanService object  will contain the ServiceID if it)
	for _, s := range *currentPlanServices {
		if s.Id == serviceId {
			return errors.Errorf("plandesigner: service '%s' already exists in Package '%s' & Plan '%s'", sourceService.Name, targetPackage.Name, targetPlan.Name)
		}
	}


	// Pre-flight checks done, good to go!

	//
	// We need to add all the existing Plan Services to the request object 'PlanDesigner' as we're calling a PUT not a PATCH
	//
	// serviceSlice = All the existing services on the Package-Plan
	//
	var serviceSlice []ServiceType

	for _, service := range *currentPlanServices {
		var endpointSlice []EndpointType
		for _, endpoint := range service.Endpoints {
			var methodSlice []MethodType
			for _, method := range endpoint.Methods {
				methodSlice = append(methodSlice, MethodType{Id: method.Id, Name: method.Name, Created: "", Updated: "", RateLimitCeiling: 0, RateLimitExempt: false, RateLimitPeriod: "day", QPSLimitCeiling: 0, QPSLimitExempt: false})
			}
			endpointSlice = append(endpointSlice, EndpointType{Id: endpoint.Id, Name: endpoint.Name, Methods: methodSlice, Created: "", Updated: "", UndefinedMethodsAllowed: true})
		}

		serviceSlice = append(serviceSlice, ServiceType{Id: service.Id, Name: service.Name, Endpoints: endpointSlice})

	}


	// Now deal with the new service we wish to add to the Package-Plan
	// Lets get down to business, create serviceType record 'sourceServiceType' that will be appended to 'serviceSlice'
	// The result will be everything we need to call the platform API to add the service


	// First we're going to have to get Endpoint Collection of the Service we're adding
	sourceEndpoints, err := endpoints.GetCollection(accessToken, &endpoints.MethodParams{ServiceId: serviceId}, &mashcli.Params{Fields: endpoints.ENDPOINTS_ALL_FIELDS})
	if err != nil {
		return err
	}

	// Now iterate through all the endpoints fetching methods and append to slice 'endpointSlice
	var endpointSlice []EndpointType
	for _, endpoint := range *sourceEndpoints {
		// Fetch all the methods for this endpoint
		sourceMethods, err := methods.GetCollection(accessToken, &methods.MethodParams{ServiceId: serviceId, EndpointId: endpoint.Id}, &mashcli.Params{Fields: methods.METHODS_ALL_FIELDS})
		if err != nil {
			return err
		}
		var methodSlice []MethodType
		for _, method := range *sourceMethods {
			methodSlice = append(methodSlice, MethodType{Id: method.Id, Name: method.Name, Created: "", Updated: "", RateLimitCeiling: 0, RateLimitExempt: false, RateLimitPeriod: "day", QPSLimitCeiling: 0, QPSLimitExempt: false})
		}
		endpointSlice = append(endpointSlice, EndpointType{Id: endpoint.Id, Name: endpoint.Name, Methods: methodSlice, Created: "", Updated: "", UndefinedMethodsAllowed: true})
	}

	// Now add the new service

	// Append newService to serviceSlice
	serviceSlice = append(serviceSlice,ServiceType{Id:sourceService.Id,Name:sourceService.Name,Endpoints: endpointSlice})

	pd := new(PlanDesigner)
	pd.Services = serviceSlice


	_, err = pd.Update(accessToken,&MethodParams{PackageId:packageId,PlanId:planId})
	if err != nil {
		return err
	}

	return nil
}

func DeleteServiceFromPackagePlan(accessToken, serviceId, packageId, planId string) error {

	// Does Package exist?
	_, err := packages.Get(accessToken, &packages.MethodParams{PackageId: packageId}, &mashcli.Params{Fields: "id,name"})
	if err != nil {
		return err
	}

	// Does Plan exist?
	_, err = plans.Get(accessToken, &plans.MethodParams{PackageId: packageId, PlanId: planId}, &mashcli.Params{Fields: "id,name"})
	if err != nil {
		return err
	}

	// Does Service exist?
	_, err = services.Get(accessToken, &services.MethodParams{ServiceId: serviceId}, &mashcli.Params{Fields: services.SERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}

	var currentPlanServices *[]planservices.PlanServices
	currentPlanServices, err = planservices.GetCollection(accessToken, &planservices.MethodParams{PackageId: packageId, PlanId: planId}, &mashcli.Params{Fields: planservices.PLANSERVICES_ALL_FIELDS})
	if err != nil {
		return err
	}


	// Pre-flight checks done, good to go!


	var serviceSlice []ServiceType

	for _, service := range *currentPlanServices {

		if service.Id != serviceId {
			var endpointSlice []EndpointType
			for _, endpoint := range service.Endpoints {
				var methodSlice []MethodType
				for _, method := range endpoint.Methods {
					methodSlice = append(methodSlice, MethodType{Id: method.Id, Name: method.Name, Created: "", Updated: "", RateLimitCeiling: 0, RateLimitExempt: false, RateLimitPeriod: "day", QPSLimitCeiling: 0, QPSLimitExempt: false})
				}
				endpointSlice = append(endpointSlice, EndpointType{Id: endpoint.Id, Name: endpoint.Name, Methods: methodSlice, Created: "", Updated: "", UndefinedMethodsAllowed: true})
			}

			serviceSlice = append(serviceSlice, ServiceType{Id: service.Id, Name: service.Name, Endpoints: endpointSlice})
		}
	}

	pd := new(PlanDesigner)
	pd.Services = serviceSlice

	_, err = pd.Update(accessToken,&MethodParams{PackageId:packageId,PlanId:planId})
	if err != nil {
		return err
	}

	return nil
}
