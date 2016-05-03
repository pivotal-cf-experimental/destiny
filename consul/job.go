package consul

import "github.com/pivotal-cf-experimental/destiny/core"

func SetJobInstanceCount(job core.Job, network core.Network, properties Properties, count int) (core.Job, Properties) {
	job.Instances = count
	for i, net := range job.Networks {
		if net.Name == network.Name {
			net.StaticIPs = network.StaticIPs(count, 0)
			properties.Consul.Agent.Servers.Lan = net.StaticIPs
		}
		job.Networks[i] = net
	}

	return job, properties
}
