# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FirewallObjectIppoolArgs', 'FirewallObjectIppool']

@pulumi.input_type
class FirewallObjectIppoolArgs:
    def __init__(__self__, *,
                 endip: pulumi.Input[str],
                 startip: pulumi.Input[str],
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallObjectIppool resource.
        :param pulumi.Input[str] endip: Final IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] startip: First IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_intf: Select an interface that will reply to ARP requests.
        :param pulumi.Input[str] arp_reply: Enable/disable replying to ARP request, default is "enable".
        :param pulumi.Input[str] associated_intf: Associated interface name.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] name: Ippool name.
        :param pulumi.Input[str] type: Ip pool type, Enum: ["overload", "one-to-one"].
        """
        pulumi.set(__self__, "endip", endip)
        pulumi.set(__self__, "startip", startip)
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if arp_intf is not None:
            pulumi.set(__self__, "arp_intf", arp_intf)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if associated_intf is not None:
            pulumi.set(__self__, "associated_intf", associated_intf)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Input[str]:
        """
        Final IPv4 address (inclusive) in the range for the address pool.
        """
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: pulumi.Input[str]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Input[str]:
        """
        First IPv4 address (inclusive) in the range for the address pool.
        """
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: pulumi.Input[str]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> Optional[pulumi.Input[str]]:
        """
        Select an interface that will reply to ARP requests.
        """
        return pulumi.get(self, "arp_intf")

    @arp_intf.setter
    def arp_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_intf", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable replying to ARP request, default is "enable".
        """
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter(name="associatedIntf")
    def associated_intf(self) -> Optional[pulumi.Input[str]]:
        """
        Associated interface name.
        """
        return pulumi.get(self, "associated_intf")

    @associated_intf.setter
    def associated_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_intf", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Ippool name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Ip pool type, Enum: ["overload", "one-to-one"].
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _FirewallObjectIppoolState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallObjectIppool resources.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_intf: Select an interface that will reply to ARP requests.
        :param pulumi.Input[str] arp_reply: Enable/disable replying to ARP request, default is "enable".
        :param pulumi.Input[str] associated_intf: Associated interface name.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] endip: Final IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] name: Ippool name.
        :param pulumi.Input[str] startip: First IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] type: Ip pool type, Enum: ["overload", "one-to-one"].
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if arp_intf is not None:
            pulumi.set(__self__, "arp_intf", arp_intf)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if associated_intf is not None:
            pulumi.set(__self__, "associated_intf", associated_intf)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if endip is not None:
            pulumi.set(__self__, "endip", endip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if startip is not None:
            pulumi.set(__self__, "startip", startip)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> Optional[pulumi.Input[str]]:
        """
        Select an interface that will reply to ARP requests.
        """
        return pulumi.get(self, "arp_intf")

    @arp_intf.setter
    def arp_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_intf", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable replying to ARP request, default is "enable".
        """
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

    @property
    @pulumi.getter(name="associatedIntf")
    def associated_intf(self) -> Optional[pulumi.Input[str]]:
        """
        Associated interface name.
        """
        return pulumi.get(self, "associated_intf")

    @associated_intf.setter
    def associated_intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_intf", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def endip(self) -> Optional[pulumi.Input[str]]:
        """
        Final IPv4 address (inclusive) in the range for the address pool.
        """
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Ippool name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def startip(self) -> Optional[pulumi.Input[str]]:
        """
        First IPv4 address (inclusive) in the range for the address pool.
        """
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Ip pool type, Enum: ["overload", "one-to-one"].
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FirewallObjectIppool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports Create/Read/Update/Delete firewall object ippool for FortiManager.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.FirewallObjectIppool("test1",
            arp_intf="any",
            arp_reply="enable",
            associated_intf="any",
            comment="test obj ippool",
            endip="1.1.10.100",
            startip="1.1.10.1",
            type="one-to-one")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_intf: Select an interface that will reply to ARP requests.
        :param pulumi.Input[str] arp_reply: Enable/disable replying to ARP request, default is "enable".
        :param pulumi.Input[str] associated_intf: Associated interface name.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] endip: Final IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] name: Ippool name.
        :param pulumi.Input[str] startip: First IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] type: Ip pool type, Enum: ["overload", "one-to-one"].
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallObjectIppoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports Create/Read/Update/Delete firewall object ippool for FortiManager.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.FirewallObjectIppool("test1",
            arp_intf="any",
            arp_reply="enable",
            associated_intf="any",
            comment="test obj ippool",
            endip="1.1.10.100",
            startip="1.1.10.1",
            type="one-to-one")
        ```

        :param str resource_name: The name of the resource.
        :param FirewallObjectIppoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallObjectIppoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 arp_intf: Optional[pulumi.Input[str]] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 associated_intf: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallObjectIppoolArgs.__new__(FirewallObjectIppoolArgs)

            __props__.__dict__["adom"] = adom
            __props__.__dict__["arp_intf"] = arp_intf
            __props__.__dict__["arp_reply"] = arp_reply
            __props__.__dict__["associated_intf"] = associated_intf
            __props__.__dict__["comment"] = comment
            if endip is None and not opts.urn:
                raise TypeError("Missing required property 'endip'")
            __props__.__dict__["endip"] = endip
            __props__.__dict__["name"] = name
            if startip is None and not opts.urn:
                raise TypeError("Missing required property 'startip'")
            __props__.__dict__["startip"] = startip
            __props__.__dict__["type"] = type
        super(FirewallObjectIppool, __self__).__init__(
            'fortios:fmg/firewallObjectIppool:FirewallObjectIppool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            arp_intf: Optional[pulumi.Input[str]] = None,
            arp_reply: Optional[pulumi.Input[str]] = None,
            associated_intf: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            endip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            startip: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FirewallObjectIppool':
        """
        Get an existing FirewallObjectIppool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: ADOM name. default is 'root'.
        :param pulumi.Input[str] arp_intf: Select an interface that will reply to ARP requests.
        :param pulumi.Input[str] arp_reply: Enable/disable replying to ARP request, default is "enable".
        :param pulumi.Input[str] associated_intf: Associated interface name.
        :param pulumi.Input[str] comment: Comments.
        :param pulumi.Input[str] endip: Final IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] name: Ippool name.
        :param pulumi.Input[str] startip: First IPv4 address (inclusive) in the range for the address pool.
        :param pulumi.Input[str] type: Ip pool type, Enum: ["overload", "one-to-one"].
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallObjectIppoolState.__new__(_FirewallObjectIppoolState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["arp_intf"] = arp_intf
        __props__.__dict__["arp_reply"] = arp_reply
        __props__.__dict__["associated_intf"] = associated_intf
        __props__.__dict__["comment"] = comment
        __props__.__dict__["endip"] = endip
        __props__.__dict__["name"] = name
        __props__.__dict__["startip"] = startip
        __props__.__dict__["type"] = type
        return FirewallObjectIppool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        """
        ADOM name. default is 'root'.
        """
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="arpIntf")
    def arp_intf(self) -> pulumi.Output[Optional[str]]:
        """
        Select an interface that will reply to ARP requests.
        """
        return pulumi.get(self, "arp_intf")

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> pulumi.Output[Optional[str]]:
        """
        Enable/disable replying to ARP request, default is "enable".
        """
        return pulumi.get(self, "arp_reply")

    @property
    @pulumi.getter(name="associatedIntf")
    def associated_intf(self) -> pulumi.Output[Optional[str]]:
        """
        Associated interface name.
        """
        return pulumi.get(self, "associated_intf")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comments.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Output[str]:
        """
        Final IPv4 address (inclusive) in the range for the address pool.
        """
        return pulumi.get(self, "endip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Ippool name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Output[str]:
        """
        First IPv4 address (inclusive) in the range for the address pool.
        """
        return pulumi.get(self, "startip")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Ip pool type, Enum: ["overload", "one-to-one"].
        """
        return pulumi.get(self, "type")

