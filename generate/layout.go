//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package main

var layout = apiInfo{
	"LoadBalancerService": {
		"addNetscalerLoadBalancer",
		"assignCertToLoadBalancer",
		"assignToGlobalLoadBalancerRule",
		"assignToLoadBalancerRule",
		"configureNetscalerLoadBalancer",
		"createGlobalLoadBalancerRule",
		"createLBHealthCheckPolicy",
		"createLBStickinessPolicy",
		"createLoadBalancer",
		"createLoadBalancerRule",
		"deleteGlobalLoadBalancerRule",
		"deleteLBHealthCheckPolicy",
		"deleteLBStickinessPolicy",
		"deleteLoadBalancer",
		"deleteLoadBalancerRule",
		"deleteNetscalerLoadBalancer",
		"deleteSslCert",
		"listGlobalLoadBalancerRules",
		"listLBHealthCheckPolicies",
		"listLBStickinessPolicies",
		"listLoadBalancerRuleInstances",
		"listLoadBalancerRules",
		"listLoadBalancers",
		"listNetscalerLoadBalancers",
		"listSslCerts",
		"removeCertFromLoadBalancer",
		"removeFromGlobalLoadBalancerRule",
		"removeFromLoadBalancerRule",
		"updateGlobalLoadBalancerRule",
		"updateLBHealthCheckPolicy",
		"updateLBStickinessPolicy",
		"updateLoadBalancer",
		"updateLoadBalancerRule",
		"uploadSslCert",
	},
	"NetworkService": {
		"addNetworkServiceProvider",
		"addOpenDaylightController",
		"createNetwork",
		"createPhysicalNetwork",
		"createServiceInstance",
		"createStorageNetworkIpRange",
		"dedicatePublicIpRange",
		"deleteNetwork",
		"deleteNetworkServiceProvider",
		"deleteOpenDaylightController",
		"deletePhysicalNetwork",
		"deleteStorageNetworkIpRange",
		"listNetscalerLoadBalancerNetworks",
		"listNetworkIsolationMethods",
		"listNetworkServiceProviders",
		"listNetworks",
		"listNiciraNvpDeviceNetworks",
		"listOpenDaylightControllers",
		"listPaloAltoFirewallNetworks",
		"listPhysicalNetworks",
		"listStorageNetworkIpRange",
		"listSupportedNetworkServices",
		"releasePublicIpRange",
		"restartNetwork",
		"updateNetwork",
		"updateNetworkServiceProvider",
		"updatePhysicalNetwork",
		"updateStorageNetworkIpRange",
		"deleteGuestNetworkIpv6Prefix",
		"createGuestNetworkIpv6Prefix",
		"listGuestNetworkIpv6Prefixes",
		"createNetworkPermissions",
		"resetNetworkPermissions",
		"listNetworkPermissions",
		"removeNetworkPermissions",
	},
	"VirtualMachineService": {
		"addNicToVirtualMachine",
		"assignVirtualMachine",
		"changeServiceForVirtualMachine",
		"cleanVMReservations",
		"deployVirtualMachine",
		"destroyVirtualMachine",
		"expungeVirtualMachine",
		"getVMPassword",
		"listVirtualMachines",
		"listVirtualMachinesMetrics",
		"migrateVirtualMachine",
		"migrateVirtualMachineWithVolume",
		"rebootVirtualMachine",
		"recoverVirtualMachine",
		"removeNicFromVirtualMachine",
		"resetPasswordForVirtualMachine",
		"restoreVirtualMachine",
		"scaleVirtualMachine",
		"startVirtualMachine",
		"stopVirtualMachine",
		"updateDefaultNicForVirtualMachine",
		"updateVirtualMachine",
		"listVirtualMachinesUsageHistory",
	},
	"VPNService": {
		"addVpnUser",
		"createRemoteAccessVpn",
		"createVpnConnection",
		"createVpnCustomerGateway",
		"createVpnGateway",
		"deleteRemoteAccessVpn",
		"deleteVpnConnection",
		"deleteVpnCustomerGateway",
		"deleteVpnGateway",
		"listRemoteAccessVpns",
		"listVpnConnections",
		"listVpnCustomerGateways",
		"listVpnGateways",
		"listVpnUsers",
		"removeVpnUser",
		"resetVpnConnection",
		"updateRemoteAccessVpn",
		"updateVpnConnection",
		"updateVpnCustomerGateway",
		"updateVpnGateway",
	},
	"FirewallService": {
		"addPaloAltoFirewall",
		"configurePaloAltoFirewall",
		"createEgressFirewallRule",
		"createFirewallRule",
		"createPortForwardingRule",
		"deleteEgressFirewallRule",
		"deleteFirewallRule",
		"deletePaloAltoFirewall",
		"deletePortForwardingRule",
		"listEgressFirewallRules",
		"listFirewallRules",
		"listPaloAltoFirewalls",
		"listPortForwardingRules",
		"updateEgressFirewallRule",
		"updateFirewallRule",
		"updatePortForwardingRule",
		"listIpv6FirewallRules",
		"createIpv6FirewallRule",
		"updateIpv6FirewallRule",
		"deleteIpv6FirewallRule",
	},
	"AutoScaleService": {
		"createAutoScalePolicy",
		"createAutoScaleVmGroup",
		"createAutoScaleVmProfile",
		"createCondition",
		"createCounter",
		"deleteAutoScalePolicy",
		"deleteAutoScaleVmGroup",
		"deleteAutoScaleVmProfile",
		"deleteCondition",
		"deleteCounter",
		"disableAutoScaleVmGroup",
		"enableAutoScaleVmGroup",
		"listAutoScalePolicies",
		"listAutoScaleVmGroups",
		"listAutoScaleVmProfiles",
		"listConditions",
		"listCounters",
		"updateAutoScalePolicy",
		"updateAutoScaleVmGroup",
		"updateAutoScaleVmProfile",
	},
	"HostService": {
		"addBaremetalHost",
		"addGloboDnsHost",
		"addHost",
		"addSecondaryStorage",
		"cancelHostMaintenance",
		"configureHAForHost",
		"enableHAForHost",
		"dedicateHost",
		"deleteHost",
		"disableOutOfBandManagementForHost",
		"enableOutOfBandManagementForHost",
		"findHostsForMigration",
		"listDedicatedHosts",
		"listHostTags",
		"listHosts",
		"listHostsMetrics",
		"prepareHostForMaintenance",
		"reconnectHost",
		"releaseDedicatedHost",
		"releaseHostReservation",
		"updateHost",
		"updateHostPassword",
	},
	"VolumeService": {
		"attachVolume",
		"createVolume",
		"deleteVolume",
		"destroyVolume",
		"detachVolume",
		"extractVolume",
		"getPathForVolume",
		"getSolidFireVolumeSize",
		"getUploadParamsForVolume",
		"getVolumeiScsiName",
		"listVolumes",
		"listVolumesMetrics",
		"migrateVolume",
		"recoverVolume",
		"resizeVolume",
		"updateVolume",
		"uploadVolume",
		"changeOfferingForVolume",
	},
	"VPCService": {
		"createPrivateGateway",
		"createStaticRoute",
		"createVPC",
		"createVPCOffering",
		"deletePrivateGateway",
		"deleteStaticRoute",
		"deleteVPC",
		"deleteVPCOffering",
		"listPrivateGateways",
		"listStaticRoutes",
		"listVPCOfferings",
		"listVPCs",
		"restartVPC",
		"updateVPC",
		"updateVPCOffering",
	},
	"TemplateService": {
		"copyTemplate",
		"createTemplate",
		"deleteTemplate",
		"extractTemplate",
		"getUploadParamsForTemplate",
		"listTemplatePermissions",
		"listTemplates",
		"prepareTemplate",
		"registerTemplate",
		"updateTemplate",
		"updateTemplatePermissions",
		"upgradeRouterTemplate",
		"listTemplateDirectDownloadCertificates",
		"provisionTemplateDirectDownloadCertificate",
	},
	"AccountService": {
		"createAccount",
		"deleteAccount",
		"disableAccount",
		"enableAccount",
		"getSolidFireAccountId",
		"listAccounts",
		"listProjectAccounts",
		"lockAccount",
		"markDefaultZoneForAccount",
		"updateAccount",
	},
	"ZoneService": {
		"createZone",
		"dedicateZone",
		"deleteZone",
		"disableOutOfBandManagementForZone",
		"enableOutOfBandManagementForZone",
		"disableHAForZone",
		"enableHAForZone",
		"listDedicatedZones",
		"listZones",
		"listZonesMetrics",
		"releaseDedicatedZone",
		"updateZone",
	},
	"UsageService": {
		"addTrafficMonitor",
		"addTrafficType",
		"deleteTrafficMonitor",
		"deleteTrafficType",
		"generateUsageRecords",
		"listTrafficMonitors",
		"listTrafficTypeImplementors",
		"listTrafficTypes",
		"listUsageRecords",
		"listUsageTypes",
		"removeRawUsageRecords",
		"updateTrafficType",
		"listUsageServerMetrics",
	},
	"SnapshotService": {
		"createSnapshot",
		"createSnapshotPolicy",
		"createVMSnapshot",
		"deleteSnapshot",
		"deleteSnapshotPolicies",
		"deleteVMSnapshot",
		"listSnapshotPolicies",
		"listSnapshots",
		"listVMSnapshot",
		"revertSnapshot",
		"revertToVMSnapshot",
		"updateSnapshotPolicy",
	},
	"UserService": {
		"createUser",
		"deleteUser",
		"disableUser",
		"enableUser",
		"getUser",
		"getUserKeys",
		"getVirtualMachineUserData",
		"listUsers",
		"lockUser",
		"registerUserKeys",
		"updateUser",
	},
	"LDAPService": {
		"addLdapConfiguration",
		"deleteLdapConfiguration",
		"importLdapUsers",
		"ldapConfig",
		"ldapCreateAccount",
		"ldapRemove",
		"linkDomainToLdap",
		"listLdapConfigurations",
		"listLdapUsers",
		"searchLdap",
	},
	"ISOService": {
		"attachIso",
		"copyIso",
		"deleteIso",
		"detachIso",
		"extractIso",
		"listIsoPermissions",
		"listIsos",
		"registerIso",
		"updateIso",
		"updateIsoPermissions",
	},
	"RouterService": {
		"changeServiceForRouter",
		"configureVirtualRouterElement",
		"createVirtualRouterElement",
		"destroyRouter",
		"listRouters",
		"listVirtualRouterElements",
		"rebootRouter",
		"startRouter",
		"stopRouter",
	},
	"ProjectService": {
		"activateProject",
		"addAccountToProject",
		"addUserToProject",
		"createProject",
		"deleteAccountFromProject",
		"deleteUserFromProject",
		"deleteProject",
		"deleteProjectInvitation",
		"listProjectInvitations",
		"listProjects",
		"suspendProject",
		"updateProject",
		"updateProjectInvitation",
		"listProjectRolePermissions",
		"createProjectRolePermission",
		"updateProjectRolePermission",
		"deleteProjectRolePermission",
	},
	"PoolService": {
		"createStoragePool",
		"deleteStoragePool",
		"findStoragePoolsForMigration",
		"listStoragePools",
		"syncStoragePool",
		"updateStoragePool",
	},
	"NetworkACLService": {
		"createNetworkACL",
		"createNetworkACLList",
		"deleteNetworkACL",
		"deleteNetworkACLList",
		"listNetworkACLLists",
		"listNetworkACLs",
		"replaceNetworkACLList",
		"updateNetworkACLItem",
		"updateNetworkACLList",
	},
	"GuestOSService": {
		"addGuestOs",
		"addGuestOsMapping",
		"listGuestOsMapping",
		"listOsCategories",
		"listOsTypes",
		"removeGuestOs",
		"removeGuestOsMapping",
		"updateGuestOs",
		"updateGuestOsMapping",
	},
	"ClusterService": {
		"addCluster",
		"dedicateCluster",
		"deleteCluster",
		"disableOutOfBandManagementForCluster",
		"enableOutOfBandManagementForCluster",
		"enableHAForCluster",
		"disableHAForCluster",
		"listClusters",
		"listClustersMetrics",
		"listDedicatedClusters",
		"releaseDedicatedCluster",
		"updateCluster",
	},
	"BaremetalService": {
		"addBaremetalDhcp",
		"addBaremetalPxeKickStartServer",
		"addBaremetalPxePingServer",
		"addBaremetalRct",
		"deleteBaremetalRct",
		"listBaremetalDhcp",
		"listBaremetalPxeServers",
		"listBaremetalRct",
		"notifyBaremetalProvisionDone",
	},
	"SystemVMService": {
		"changeServiceForSystemVm",
		"destroySystemVm",
		"listSystemVms",
		"migrateSystemVm",
		"rebootSystemVm",
		"scaleSystemVm",
		"startSystemVm",
		"stopSystemVm",
		"patchSystemVm",
	},
	"RoleService": {
		"createRole",
		"createRolePermission",
		"deleteRole",
		"deleteRolePermission",
		"importRole",
		"listRolePermissions",
		"listRoles",
		"updateRole",
		"updateRolePermission",
	},
	"ImageStoreService": {
		"addImageStore",
		"addImageStoreS3",
		"createSecondaryStagingStore",
		"deleteImageStore",
		"deleteSecondaryStagingStore",
		"listImageStores",
		"listSecondaryStagingStores",
		"updateCloudToUseObjectStore",
	},
	"AuthenticationService": {
		"login",
		"logout",
	},
	"SecurityGroupService": {
		"authorizeSecurityGroupEgress",
		"authorizeSecurityGroupIngress",
		"createSecurityGroup",
		"deleteSecurityGroup",
		"listSecurityGroups",
		"revokeSecurityGroupEgress",
		"revokeSecurityGroupIngress",
	},
	"QuotaService": {
		"quotaIsEnabled",
	},
	"PodService": {
		"createPod",
		"dedicatePod",
		"deletePod",
		"listDedicatedPods",
		"listPods",
		"releaseDedicatedPod",
		"updatePod",
	},
	"VLANService": {
		"createVlanIpRange",
		"dedicateGuestVlanRange",
		"deleteVlanIpRange",
		"listDedicatedGuestVlanRanges",
		"listVlanIpRanges",
		"releaseDedicatedGuestVlanRange",
		"listGuestVlans",
	},
	"UCSService": {
		"addUcsManager",
		"associateUcsProfileToBlade",
		"deleteUcsManager",
		"listUcsBlades",
		"listUcsManagers",
		"listUcsProfiles",
	},
	"InternalLBService": {
		"configureInternalLoadBalancerElement",
		"createInternalLoadBalancerElement",
		"listInternalLoadBalancerElements",
		"listInternalLoadBalancerVMs",
		"startInternalLoadBalancerVM",
		"stopInternalLoadBalancerVM",
	},
	"SSHService": {
		"createSSHKeyPair",
		"deleteSSHKeyPair",
		"listSSHKeyPairs",
		"registerSSHKeyPair",
		"resetSSHKeyForVirtualMachine",
	},
	"NATService": {
		"createIpForwardingRule",
		"deleteIpForwardingRule",
		"disableStaticNat",
		"enableStaticNat",
		"listIpForwardingRules",
	},
	"LimitService": {
		"getApiLimit",
		"listResourceLimits",
		"resetApiLimit",
		"updateResourceCount",
		"updateResourceLimit",
	},
	"DomainService": {
		"createDomain",
		"deleteDomain",
		"listDomainChildren",
		"listDomains",
		"updateDomain",
	},
	"AffinityGroupService": {
		"createAffinityGroup",
		"deleteAffinityGroup",
		"listAffinityGroupTypes",
		"listAffinityGroups",
		"updateVMAffinityGroup",
	},
	"VMGroupService": {
		"createInstanceGroup",
		"deleteInstanceGroup",
		"listInstanceGroups",
		"updateInstanceGroup",
	},
	"ServiceOfferingService": {
		"createServiceOffering",
		"deleteServiceOffering",
		"listServiceOfferings",
		"updateServiceOffering",
	},
	"ResourcetagsService": {
		"createTags",
		"deleteTags",
		"listStorageTags",
		"listTags",
	},
	"ResourcemetadataService": {
		"addResourceDetail",
		"getVolumeSnapshotDetails",
		"listResourceDetails",
		"removeResourceDetail",
	},
	"RegionService": {
		"addRegion",
		"listRegions",
		"removeRegion",
		"updateRegion",
	},
	"NicService": {
		"addIpToNic",
		"listNics",
		"removeIpFromNic",
		"updateVmNicIp",
	},
	"NetworkOfferingService": {
		"createNetworkOffering",
		"deleteNetworkOffering",
		"listNetworkOfferings",
		"updateNetworkOffering",
	},
	"EventService": {
		"archiveEvents",
		"deleteEvents",
		"listEventTypes",
		"listEvents",
	},
	"DiskOfferingService": {
		"createDiskOffering",
		"deleteDiskOffering",
		"listDiskOfferings",
		"updateDiskOffering",
	},
	"ConfigurationService": {
		"listCapabilities",
		"listConfigurations",
		"listDeploymentPlanners",
		"updateConfiguration",
		"resetConfiguration",
	},
	"BrocadeVCSService": {
		"addBrocadeVcsDevice",
		"deleteBrocadeVcsDevice",
		"listBrocadeVcsDeviceNetworks",
		"listBrocadeVcsDevices",
	},
	"AlertService": {
		"archiveAlerts",
		"deleteAlerts",
		"generateAlert",
		"listAlerts",
	},
	"AddressService": {
		"associateIpAddress",
		"disassociateIpAddress",
		"listPublicIpAddresses",
		"updateIpAddress",
		"releaseIpAddress",
	},
	"StoragePoolService": {
		"cancelStorageMaintenance",
		"enableStorageMaintenance",
		"listStorageProviders",
	},
	"PortableIPService": {
		"createPortableIpRange",
		"deletePortableIpRange",
		"listPortableIpRanges",
	},
	"OutofbandManagementService": {
		"changeOutOfBandManagementPassword",
		"configureOutOfBandManagement",
		"issueOutOfBandManagementPowerAction",
	},
	"NiciraNVPService": {
		"addNiciraNvpDevice",
		"deleteNiciraNvpDevice",
		"listNiciraNvpDevices",
	},
	"NetworkDeviceService": {
		"addNetworkDevice",
		"deleteNetworkDevice",
		"listNetworkDevice",
	},
	"HypervisorService": {
		"listHypervisorCapabilities",
		"listHypervisors",
		"updateHypervisorCapabilities",
	},
	"BigSwitchBCFService": {
		"addBigSwitchBcfDevice",
		"deleteBigSwitchBcfDevice",
		"listBigSwitchBcfDevices",
	},
	"SwiftService": {
		"addSwift",
		"listSwifts",
	},
	"OvsElementService": {
		"configureOvsElement",
		"listOvsElements",
	},
	"AsyncjobService": {
		"listAsyncJobs",
		"queryAsyncJobResult",
	},
	"StratosphereSSPService": {
		"addStratosphereSsp",
		"deleteStratosphereSsp",
	},
	"SystemCapacityService": {
		"listCapacity",
	},
	"CloudIdentifierService": {
		"getCloudIdentifier",
	},
	"CertificateService": {
		"uploadCustomCertificate",
	},
	"APIDiscoveryService": {
		"listApis",
	},
	"AnnotationService": {
		"addAnnotation",
		"listAnnotations",
		"removeAnnotation",
		"updateAnnotationVisibility",
	},
	"KubernetesService": {
		"addKubernetesSupportedVersion",
		"createKubernetesCluster",
		"deleteKubernetesCluster",
		"deleteKubernetesSupportedVersion",
		"getKubernetesClusterConfig",
		"listKubernetesClusters",
		"listKubernetesSupportedVersions",
		"scaleKubernetesCluster",
		"startKubernetesCluster",
		"stopKubernetesCluster",
		"updateKubernetesSupportedVersion",
		"upgradeKubernetesCluster",
	},
	"InfrastructureUsageService": {
		"listManagementServersMetrics",
		"listDbMetrics",
	},
}
