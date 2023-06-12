# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetAddressResult',
    'AwaitableGetAddressResult',
    'get_address',
    'get_address_output',
]

@pulumi.output_type
class GetAddressResult:
    """
    A collection of values returned by getAddress.
    """
    def __init__(__self__, allow_routing=None, associated_interface=None, cache_ttl=None, clearpass_spt=None, color=None, comment=None, country=None, end_ip=None, end_mac=None, epg_name=None, fabric_object=None, filter=None, fqdn=None, fsso_groups=None, id=None, interface=None, lists=None, macaddrs=None, name=None, node_ip_only=None, obj_id=None, obj_tag=None, obj_type=None, organization=None, policy_group=None, sdn=None, sdn_addr_type=None, sdn_tag=None, start_ip=None, start_mac=None, sub_type=None, subnet=None, subnet_name=None, tag_detection_level=None, tag_type=None, taggings=None, tenant=None, type=None, uuid=None, vdomparam=None, visibility=None, wildcard=None, wildcard_fqdn=None):
        if allow_routing and not isinstance(allow_routing, str):
            raise TypeError("Expected argument 'allow_routing' to be a str")
        pulumi.set(__self__, "allow_routing", allow_routing)
        if associated_interface and not isinstance(associated_interface, str):
            raise TypeError("Expected argument 'associated_interface' to be a str")
        pulumi.set(__self__, "associated_interface", associated_interface)
        if cache_ttl and not isinstance(cache_ttl, int):
            raise TypeError("Expected argument 'cache_ttl' to be a int")
        pulumi.set(__self__, "cache_ttl", cache_ttl)
        if clearpass_spt and not isinstance(clearpass_spt, str):
            raise TypeError("Expected argument 'clearpass_spt' to be a str")
        pulumi.set(__self__, "clearpass_spt", clearpass_spt)
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if country and not isinstance(country, str):
            raise TypeError("Expected argument 'country' to be a str")
        pulumi.set(__self__, "country", country)
        if end_ip and not isinstance(end_ip, str):
            raise TypeError("Expected argument 'end_ip' to be a str")
        pulumi.set(__self__, "end_ip", end_ip)
        if end_mac and not isinstance(end_mac, str):
            raise TypeError("Expected argument 'end_mac' to be a str")
        pulumi.set(__self__, "end_mac", end_mac)
        if epg_name and not isinstance(epg_name, str):
            raise TypeError("Expected argument 'epg_name' to be a str")
        pulumi.set(__self__, "epg_name", epg_name)
        if fabric_object and not isinstance(fabric_object, str):
            raise TypeError("Expected argument 'fabric_object' to be a str")
        pulumi.set(__self__, "fabric_object", fabric_object)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if fsso_groups and not isinstance(fsso_groups, list):
            raise TypeError("Expected argument 'fsso_groups' to be a list")
        pulumi.set(__self__, "fsso_groups", fsso_groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if macaddrs and not isinstance(macaddrs, list):
            raise TypeError("Expected argument 'macaddrs' to be a list")
        pulumi.set(__self__, "macaddrs", macaddrs)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_ip_only and not isinstance(node_ip_only, str):
            raise TypeError("Expected argument 'node_ip_only' to be a str")
        pulumi.set(__self__, "node_ip_only", node_ip_only)
        if obj_id and not isinstance(obj_id, str):
            raise TypeError("Expected argument 'obj_id' to be a str")
        pulumi.set(__self__, "obj_id", obj_id)
        if obj_tag and not isinstance(obj_tag, str):
            raise TypeError("Expected argument 'obj_tag' to be a str")
        pulumi.set(__self__, "obj_tag", obj_tag)
        if obj_type and not isinstance(obj_type, str):
            raise TypeError("Expected argument 'obj_type' to be a str")
        pulumi.set(__self__, "obj_type", obj_type)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if policy_group and not isinstance(policy_group, str):
            raise TypeError("Expected argument 'policy_group' to be a str")
        pulumi.set(__self__, "policy_group", policy_group)
        if sdn and not isinstance(sdn, str):
            raise TypeError("Expected argument 'sdn' to be a str")
        pulumi.set(__self__, "sdn", sdn)
        if sdn_addr_type and not isinstance(sdn_addr_type, str):
            raise TypeError("Expected argument 'sdn_addr_type' to be a str")
        pulumi.set(__self__, "sdn_addr_type", sdn_addr_type)
        if sdn_tag and not isinstance(sdn_tag, str):
            raise TypeError("Expected argument 'sdn_tag' to be a str")
        pulumi.set(__self__, "sdn_tag", sdn_tag)
        if start_ip and not isinstance(start_ip, str):
            raise TypeError("Expected argument 'start_ip' to be a str")
        pulumi.set(__self__, "start_ip", start_ip)
        if start_mac and not isinstance(start_mac, str):
            raise TypeError("Expected argument 'start_mac' to be a str")
        pulumi.set(__self__, "start_mac", start_mac)
        if sub_type and not isinstance(sub_type, str):
            raise TypeError("Expected argument 'sub_type' to be a str")
        pulumi.set(__self__, "sub_type", sub_type)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if subnet_name and not isinstance(subnet_name, str):
            raise TypeError("Expected argument 'subnet_name' to be a str")
        pulumi.set(__self__, "subnet_name", subnet_name)
        if tag_detection_level and not isinstance(tag_detection_level, str):
            raise TypeError("Expected argument 'tag_detection_level' to be a str")
        pulumi.set(__self__, "tag_detection_level", tag_detection_level)
        if tag_type and not isinstance(tag_type, str):
            raise TypeError("Expected argument 'tag_type' to be a str")
        pulumi.set(__self__, "tag_type", tag_type)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if tenant and not isinstance(tenant, str):
            raise TypeError("Expected argument 'tenant' to be a str")
        pulumi.set(__self__, "tenant", tenant)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)
        if wildcard and not isinstance(wildcard, str):
            raise TypeError("Expected argument 'wildcard' to be a str")
        pulumi.set(__self__, "wildcard", wildcard)
        if wildcard_fqdn and not isinstance(wildcard_fqdn, str):
            raise TypeError("Expected argument 'wildcard_fqdn' to be a str")
        pulumi.set(__self__, "wildcard_fqdn", wildcard_fqdn)

    @property
    @pulumi.getter(name="allowRouting")
    def allow_routing(self) -> str:
        """
        Enable/disable use of this address in the static route configuration.
        """
        return pulumi.get(self, "allow_routing")

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> str:
        """
        Network interface associated with address.
        """
        return pulumi.get(self, "associated_interface")

    @property
    @pulumi.getter(name="cacheTtl")
    def cache_ttl(self) -> int:
        """
        Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
        """
        return pulumi.get(self, "cache_ttl")

    @property
    @pulumi.getter(name="clearpassSpt")
    def clearpass_spt(self) -> str:
        """
        SPT (System Posture Token) value.
        """
        return pulumi.get(self, "clearpass_spt")

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def country(self) -> str:
        """
        IP addresses associated to a specific country.
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> str:
        """
        Final IP address (inclusive) in the range for the address.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter(name="endMac")
    def end_mac(self) -> str:
        """
        Last MAC address in the range.
        """
        return pulumi.get(self, "end_mac")

    @property
    @pulumi.getter(name="epgName")
    def epg_name(self) -> str:
        """
        Endpoint group name.
        """
        return pulumi.get(self, "epg_name")

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> str:
        """
        Security Fabric global object setting.
        """
        return pulumi.get(self, "fabric_object")

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        Match criteria filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        Fully Qualified Domain Name address.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="fssoGroups")
    def fsso_groups(self) -> Sequence['outputs.GetAddressFssoGroupResult']:
        """
        FSSO group(s). The structure of `fsso_group` block is documented below.
        """
        return pulumi.get(self, "fsso_groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Name of interface whose IP address is to be used.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetAddressListResult']:
        """
        IP address list. The structure of `list` block is documented below.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def macaddrs(self) -> Sequence['outputs.GetAddressMacaddrResult']:
        """
        MAC address ranges <start>[-<end>] separated by space.
        """
        return pulumi.get(self, "macaddrs")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tag name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeIpOnly")
    def node_ip_only(self) -> str:
        """
        Enable/disable collection of node addresses only in Kubernetes.
        """
        return pulumi.get(self, "node_ip_only")

    @property
    @pulumi.getter(name="objId")
    def obj_id(self) -> str:
        """
        Object ID for NSX.
        """
        return pulumi.get(self, "obj_id")

    @property
    @pulumi.getter(name="objTag")
    def obj_tag(self) -> str:
        """
        Tag of dynamic address object.
        """
        return pulumi.get(self, "obj_tag")

    @property
    @pulumi.getter(name="objType")
    def obj_type(self) -> str:
        """
        Object type.
        """
        return pulumi.get(self, "obj_type")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        Organization domain name (Syntax: organization/domain).
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="policyGroup")
    def policy_group(self) -> str:
        """
        Policy group name.
        """
        return pulumi.get(self, "policy_group")

    @property
    @pulumi.getter
    def sdn(self) -> str:
        """
        SDN.
        """
        return pulumi.get(self, "sdn")

    @property
    @pulumi.getter(name="sdnAddrType")
    def sdn_addr_type(self) -> str:
        """
        Type of addresses to collect.
        """
        return pulumi.get(self, "sdn_addr_type")

    @property
    @pulumi.getter(name="sdnTag")
    def sdn_tag(self) -> str:
        """
        SDN Tag.
        """
        return pulumi.get(self, "sdn_tag")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> str:
        """
        First IP address (inclusive) in the range for the address.
        """
        return pulumi.get(self, "start_ip")

    @property
    @pulumi.getter(name="startMac")
    def start_mac(self) -> str:
        """
        First MAC address in the range.
        """
        return pulumi.get(self, "start_mac")

    @property
    @pulumi.getter(name="subType")
    def sub_type(self) -> str:
        """
        Sub-type of address.
        """
        return pulumi.get(self, "sub_type")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        """
        IP address and subnet mask of address.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetName")
    def subnet_name(self) -> str:
        """
        Subnet name.
        """
        return pulumi.get(self, "subnet_name")

    @property
    @pulumi.getter(name="tagDetectionLevel")
    def tag_detection_level(self) -> str:
        """
        Tag detection level of dynamic address object.
        """
        return pulumi.get(self, "tag_detection_level")

    @property
    @pulumi.getter(name="tagType")
    def tag_type(self) -> str:
        """
        Tag type of dynamic address object.
        """
        return pulumi.get(self, "tag_type")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetAddressTaggingResult']:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def tenant(self) -> str:
        """
        Tenant.
        """
        return pulumi.get(self, "tenant")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of address.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Enable/disable address visibility in the GUI.
        """
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter
    def wildcard(self) -> str:
        """
        IP address and wildcard netmask.
        """
        return pulumi.get(self, "wildcard")

    @property
    @pulumi.getter(name="wildcardFqdn")
    def wildcard_fqdn(self) -> str:
        """
        Fully Qualified Domain Name with wildcard characters.
        """
        return pulumi.get(self, "wildcard_fqdn")


class AwaitableGetAddressResult(GetAddressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressResult(
            allow_routing=self.allow_routing,
            associated_interface=self.associated_interface,
            cache_ttl=self.cache_ttl,
            clearpass_spt=self.clearpass_spt,
            color=self.color,
            comment=self.comment,
            country=self.country,
            end_ip=self.end_ip,
            end_mac=self.end_mac,
            epg_name=self.epg_name,
            fabric_object=self.fabric_object,
            filter=self.filter,
            fqdn=self.fqdn,
            fsso_groups=self.fsso_groups,
            id=self.id,
            interface=self.interface,
            lists=self.lists,
            macaddrs=self.macaddrs,
            name=self.name,
            node_ip_only=self.node_ip_only,
            obj_id=self.obj_id,
            obj_tag=self.obj_tag,
            obj_type=self.obj_type,
            organization=self.organization,
            policy_group=self.policy_group,
            sdn=self.sdn,
            sdn_addr_type=self.sdn_addr_type,
            sdn_tag=self.sdn_tag,
            start_ip=self.start_ip,
            start_mac=self.start_mac,
            sub_type=self.sub_type,
            subnet=self.subnet,
            subnet_name=self.subnet_name,
            tag_detection_level=self.tag_detection_level,
            tag_type=self.tag_type,
            taggings=self.taggings,
            tenant=self.tenant,
            type=self.type,
            uuid=self.uuid,
            vdomparam=self.vdomparam,
            visibility=self.visibility,
            wildcard=self.wildcard,
            wildcard_fqdn=self.wildcard_fqdn)


def get_address(name: Optional[str] = None,
                vdomparam: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressResult:
    """
    Use this data source to get information on an fortios firewall address


    :param str name: Specify the name of the desired firewall address.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/getAddress:getAddress', __args__, opts=opts, typ=GetAddressResult).value

    return AwaitableGetAddressResult(
        allow_routing=__ret__.allow_routing,
        associated_interface=__ret__.associated_interface,
        cache_ttl=__ret__.cache_ttl,
        clearpass_spt=__ret__.clearpass_spt,
        color=__ret__.color,
        comment=__ret__.comment,
        country=__ret__.country,
        end_ip=__ret__.end_ip,
        end_mac=__ret__.end_mac,
        epg_name=__ret__.epg_name,
        fabric_object=__ret__.fabric_object,
        filter=__ret__.filter,
        fqdn=__ret__.fqdn,
        fsso_groups=__ret__.fsso_groups,
        id=__ret__.id,
        interface=__ret__.interface,
        lists=__ret__.lists,
        macaddrs=__ret__.macaddrs,
        name=__ret__.name,
        node_ip_only=__ret__.node_ip_only,
        obj_id=__ret__.obj_id,
        obj_tag=__ret__.obj_tag,
        obj_type=__ret__.obj_type,
        organization=__ret__.organization,
        policy_group=__ret__.policy_group,
        sdn=__ret__.sdn,
        sdn_addr_type=__ret__.sdn_addr_type,
        sdn_tag=__ret__.sdn_tag,
        start_ip=__ret__.start_ip,
        start_mac=__ret__.start_mac,
        sub_type=__ret__.sub_type,
        subnet=__ret__.subnet,
        subnet_name=__ret__.subnet_name,
        tag_detection_level=__ret__.tag_detection_level,
        tag_type=__ret__.tag_type,
        taggings=__ret__.taggings,
        tenant=__ret__.tenant,
        type=__ret__.type,
        uuid=__ret__.uuid,
        vdomparam=__ret__.vdomparam,
        visibility=__ret__.visibility,
        wildcard=__ret__.wildcard,
        wildcard_fqdn=__ret__.wildcard_fqdn)


@_utilities.lift_output_func(get_address)
def get_address_output(name: Optional[pulumi.Input[str]] = None,
                       vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAddressResult]:
    """
    Use this data source to get information on an fortios firewall address


    :param str name: Specify the name of the desired firewall address.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...