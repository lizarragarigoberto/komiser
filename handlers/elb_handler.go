package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) ElasticLoadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("elb")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeElasticLoadBalancer(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "elasticloadbalancing:DescribeLoadBalancers is missing")
		} else {
			handler.cache.Set("elb", response, cache.DefaultExpiration)
			respondWithJSON(w, 200, response)
		}
	}
}
