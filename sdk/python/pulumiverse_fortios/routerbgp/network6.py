# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Network6Args', 'Network6']

@pulumi.input_type
class Network6Args:
    def __init__(__self__, *,
                 backdoor: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 network_import_check: Optional[pulumi.Input[str]] = None,
                 prefix6: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Network6 resource.
        :param pulumi.Input[str] backdoor: Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] network_import_check: Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        :param pulumi.Input[str] prefix6: Network IPv6 prefix.
        :param pulumi.Input[str] route_map: Route map to modify generated route.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if backdoor is not None:
            pulumi.set(__self__, "backdoor", backdoor)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if network_import_check is not None:
            pulumi.set(__self__, "network_import_check", network_import_check)
        if prefix6 is not None:
            pulumi.set(__self__, "prefix6", prefix6)
        if route_map is not None:
            pulumi.set(__self__, "route_map", route_map)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def backdoor(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "backdoor")

    @backdoor.setter
    def backdoor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backdoor", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="networkImportCheck")
    def network_import_check(self) -> Optional[pulumi.Input[str]]:
        """
        Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        """
        return pulumi.get(self, "network_import_check")

    @network_import_check.setter
    def network_import_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_import_check", value)

    @property
    @pulumi.getter
    def prefix6(self) -> Optional[pulumi.Input[str]]:
        """
        Network IPv6 prefix.
        """
        return pulumi.get(self, "prefix6")

    @prefix6.setter
    def prefix6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix6", value)

    @property
    @pulumi.getter(name="routeMap")
    def route_map(self) -> Optional[pulumi.Input[str]]:
        """
        Route map to modify generated route.
        """
        return pulumi.get(self, "route_map")

    @route_map.setter
    def route_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_map", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _Network6State:
    def __init__(__self__, *,
                 backdoor: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 network_import_check: Optional[pulumi.Input[str]] = None,
                 prefix6: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Network6 resources.
        :param pulumi.Input[str] backdoor: Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] network_import_check: Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        :param pulumi.Input[str] prefix6: Network IPv6 prefix.
        :param pulumi.Input[str] route_map: Route map to modify generated route.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if backdoor is not None:
            pulumi.set(__self__, "backdoor", backdoor)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if network_import_check is not None:
            pulumi.set(__self__, "network_import_check", network_import_check)
        if prefix6 is not None:
            pulumi.set(__self__, "prefix6", prefix6)
        if route_map is not None:
            pulumi.set(__self__, "route_map", route_map)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def backdoor(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "backdoor")

    @backdoor.setter
    def backdoor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backdoor", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="networkImportCheck")
    def network_import_check(self) -> Optional[pulumi.Input[str]]:
        """
        Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        """
        return pulumi.get(self, "network_import_check")

    @network_import_check.setter
    def network_import_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_import_check", value)

    @property
    @pulumi.getter
    def prefix6(self) -> Optional[pulumi.Input[str]]:
        """
        Network IPv6 prefix.
        """
        return pulumi.get(self, "prefix6")

    @prefix6.setter
    def prefix6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix6", value)

    @property
    @pulumi.getter(name="routeMap")
    def route_map(self) -> Optional[pulumi.Input[str]]:
        """
        Route map to modify generated route.
        """
        return pulumi.get(self, "route_map")

    @route_map.setter
    def route_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_map", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Network6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backdoor: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 network_import_check: Optional[pulumi.Input[str]] = None,
                 prefix6: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        BGP IPv6 network table.

        > The provider supports the definition of Network6 in Router Bgp `router.Bgp`, and also allows the definition of separate Network6 resources `routerbgp.Network6`, but do not use a `router.Bgp` with in-line Network6 in conjunction with any `routerbgp.Network6` resources, otherwise conflicts and overwrite will occur.

        ## Import

        Routerbgp Network6 can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:routerbgp/network6:Network6 labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:routerbgp/network6:Network6 labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backdoor: Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] network_import_check: Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        :param pulumi.Input[str] prefix6: Network IPv6 prefix.
        :param pulumi.Input[str] route_map: Route map to modify generated route.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[Network6Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        BGP IPv6 network table.

        > The provider supports the definition of Network6 in Router Bgp `router.Bgp`, and also allows the definition of separate Network6 resources `routerbgp.Network6`, but do not use a `router.Bgp` with in-line Network6 in conjunction with any `routerbgp.Network6` resources, otherwise conflicts and overwrite will occur.

        ## Import

        Routerbgp Network6 can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:routerbgp/network6:Network6 labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:routerbgp/network6:Network6 labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param Network6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Network6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backdoor: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 network_import_check: Optional[pulumi.Input[str]] = None,
                 prefix6: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Network6Args.__new__(Network6Args)

            __props__.__dict__["backdoor"] = backdoor
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["network_import_check"] = network_import_check
            __props__.__dict__["prefix6"] = prefix6
            __props__.__dict__["route_map"] = route_map
            __props__.__dict__["vdomparam"] = vdomparam
        super(Network6, __self__).__init__(
            'fortios:routerbgp/network6:Network6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backdoor: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            network_import_check: Optional[pulumi.Input[str]] = None,
            prefix6: Optional[pulumi.Input[str]] = None,
            route_map: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Network6':
        """
        Get an existing Network6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backdoor: Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] network_import_check: Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        :param pulumi.Input[str] prefix6: Network IPv6 prefix.
        :param pulumi.Input[str] route_map: Route map to modify generated route.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Network6State.__new__(_Network6State)

        __props__.__dict__["backdoor"] = backdoor
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["network_import_check"] = network_import_check
        __props__.__dict__["prefix6"] = prefix6
        __props__.__dict__["route_map"] = route_map
        __props__.__dict__["vdomparam"] = vdomparam
        return Network6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backdoor(self) -> pulumi.Output[str]:
        """
        Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "backdoor")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="networkImportCheck")
    def network_import_check(self) -> pulumi.Output[str]:
        """
        Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        """
        return pulumi.get(self, "network_import_check")

    @property
    @pulumi.getter
    def prefix6(self) -> pulumi.Output[str]:
        """
        Network IPv6 prefix.
        """
        return pulumi.get(self, "prefix6")

    @property
    @pulumi.getter(name="routeMap")
    def route_map(self) -> pulumi.Output[str]:
        """
        Route map to modify generated route.
        """
        return pulumi.get(self, "route_map")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

