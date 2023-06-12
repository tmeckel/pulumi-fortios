# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StormcontrolpolicyArgs', 'Stormcontrolpolicy']

@pulumi.input_type
class StormcontrolpolicyArgs:
    def __init__(__self__, *,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Stormcontrolpolicy resource.
        :param pulumi.Input[str] broadcast: Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description of the storm control policy.
        :param pulumi.Input[str] name: Storm control policy name.
        :param pulumi.Input[int] rate: Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        :param pulumi.Input[str] storm_control_mode: Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        :param pulumi.Input[str] unknown_multicast: Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] unknown_unicast: Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if broadcast is not None:
            pulumi.set(__self__, "broadcast", broadcast)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate is not None:
            pulumi.set(__self__, "rate", rate)
        if storm_control_mode is not None:
            pulumi.set(__self__, "storm_control_mode", storm_control_mode)
        if unknown_multicast is not None:
            pulumi.set(__self__, "unknown_multicast", unknown_multicast)
        if unknown_unicast is not None:
            pulumi.set(__self__, "unknown_unicast", unknown_unicast)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def broadcast(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "broadcast")

    @broadcast.setter
    def broadcast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "broadcast", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the storm control policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Storm control policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rate(self) -> Optional[pulumi.Input[int]]:
        """
        Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        """
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter(name="stormControlMode")
    def storm_control_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        """
        return pulumi.get(self, "storm_control_mode")

    @storm_control_mode.setter
    def storm_control_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storm_control_mode", value)

    @property
    @pulumi.getter(name="unknownMulticast")
    def unknown_multicast(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "unknown_multicast")

    @unknown_multicast.setter
    def unknown_multicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_multicast", value)

    @property
    @pulumi.getter(name="unknownUnicast")
    def unknown_unicast(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "unknown_unicast")

    @unknown_unicast.setter
    def unknown_unicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_unicast", value)

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
class _StormcontrolpolicyState:
    def __init__(__self__, *,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Stormcontrolpolicy resources.
        :param pulumi.Input[str] broadcast: Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description of the storm control policy.
        :param pulumi.Input[str] name: Storm control policy name.
        :param pulumi.Input[int] rate: Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        :param pulumi.Input[str] storm_control_mode: Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        :param pulumi.Input[str] unknown_multicast: Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] unknown_unicast: Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if broadcast is not None:
            pulumi.set(__self__, "broadcast", broadcast)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate is not None:
            pulumi.set(__self__, "rate", rate)
        if storm_control_mode is not None:
            pulumi.set(__self__, "storm_control_mode", storm_control_mode)
        if unknown_multicast is not None:
            pulumi.set(__self__, "unknown_multicast", unknown_multicast)
        if unknown_unicast is not None:
            pulumi.set(__self__, "unknown_unicast", unknown_unicast)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def broadcast(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "broadcast")

    @broadcast.setter
    def broadcast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "broadcast", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the storm control policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Storm control policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rate(self) -> Optional[pulumi.Input[int]]:
        """
        Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        """
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter(name="stormControlMode")
    def storm_control_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        """
        return pulumi.get(self, "storm_control_mode")

    @storm_control_mode.setter
    def storm_control_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storm_control_mode", value)

    @property
    @pulumi.getter(name="unknownMulticast")
    def unknown_multicast(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "unknown_multicast")

    @unknown_multicast.setter
    def unknown_multicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_multicast", value)

    @property
    @pulumi.getter(name="unknownUnicast")
    def unknown_unicast(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "unknown_unicast")

    @unknown_unicast.setter
    def unknown_unicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_unicast", value)

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


class Stormcontrolpolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSwitch storm control policy to be applied on managed-switch ports. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController StormControlPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:switchcontroller/stormcontrolpolicy:Stormcontrolpolicy labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:switchcontroller/stormcontrolpolicy:Stormcontrolpolicy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] broadcast: Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description of the storm control policy.
        :param pulumi.Input[str] name: Storm control policy name.
        :param pulumi.Input[int] rate: Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        :param pulumi.Input[str] storm_control_mode: Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        :param pulumi.Input[str] unknown_multicast: Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] unknown_unicast: Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StormcontrolpolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSwitch storm control policy to be applied on managed-switch ports. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController StormControlPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:switchcontroller/stormcontrolpolicy:Stormcontrolpolicy labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:switchcontroller/stormcontrolpolicy:Stormcontrolpolicy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param StormcontrolpolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StormcontrolpolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StormcontrolpolicyArgs.__new__(StormcontrolpolicyArgs)

            __props__.__dict__["broadcast"] = broadcast
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["rate"] = rate
            __props__.__dict__["storm_control_mode"] = storm_control_mode
            __props__.__dict__["unknown_multicast"] = unknown_multicast
            __props__.__dict__["unknown_unicast"] = unknown_unicast
            __props__.__dict__["vdomparam"] = vdomparam
        super(Stormcontrolpolicy, __self__).__init__(
            'fortios:switchcontroller/stormcontrolpolicy:Stormcontrolpolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            broadcast: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rate: Optional[pulumi.Input[int]] = None,
            storm_control_mode: Optional[pulumi.Input[str]] = None,
            unknown_multicast: Optional[pulumi.Input[str]] = None,
            unknown_unicast: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Stormcontrolpolicy':
        """
        Get an existing Stormcontrolpolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] broadcast: Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description of the storm control policy.
        :param pulumi.Input[str] name: Storm control policy name.
        :param pulumi.Input[int] rate: Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        :param pulumi.Input[str] storm_control_mode: Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        :param pulumi.Input[str] unknown_multicast: Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] unknown_unicast: Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StormcontrolpolicyState.__new__(_StormcontrolpolicyState)

        __props__.__dict__["broadcast"] = broadcast
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["rate"] = rate
        __props__.__dict__["storm_control_mode"] = storm_control_mode
        __props__.__dict__["unknown_multicast"] = unknown_multicast
        __props__.__dict__["unknown_unicast"] = unknown_unicast
        __props__.__dict__["vdomparam"] = vdomparam
        return Stormcontrolpolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def broadcast(self) -> pulumi.Output[str]:
        """
        Enable/disable storm control to drop/allow broadcast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "broadcast")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the storm control policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Storm control policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Output[int]:
        """
        Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
        """
        return pulumi.get(self, "rate")

    @property
    @pulumi.getter(name="stormControlMode")
    def storm_control_mode(self) -> pulumi.Output[str]:
        """
        Set Storm control mode. Valid values: `global`, `override`, `disabled`.
        """
        return pulumi.get(self, "storm_control_mode")

    @property
    @pulumi.getter(name="unknownMulticast")
    def unknown_multicast(self) -> pulumi.Output[str]:
        """
        Enable/disable storm control to drop/allow unknown multicast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "unknown_multicast")

    @property
    @pulumi.getter(name="unknownUnicast")
    def unknown_unicast(self) -> pulumi.Output[str]:
        """
        Enable/disable storm control to drop/allow unknown unicast traffic in override mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "unknown_unicast")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
