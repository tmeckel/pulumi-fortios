# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FortigateArgs', 'Fortigate']

@pulumi.input_type
class FortigateArgs:
    def __init__(__self__, *,
                 authorized: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fortigate resource.
        :param pulumi.Input[str] authorized: Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] device_id: device-id
        :param pulumi.Input[str] fosid: FortiGate serial number.
        :param pulumi.Input[str] hostname: FortiGate hostname.
        :param pulumi.Input[str] name: FortiGate entry name.
        :param pulumi.Input[str] profile: FortiGate profile configuration.
        :param pulumi.Input[int] vdom: VDOM.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authorized is not None:
            pulumi.set(__self__, "authorized", authorized)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authorized(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[int]]:
        """
        device-id
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate serial number.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate hostname.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate profile configuration.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[int]]:
        """
        VDOM.
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vdom", value)

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
class _FortigateState:
    def __init__(__self__, *,
                 authorized: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Fortigate resources.
        :param pulumi.Input[str] authorized: Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] device_id: device-id
        :param pulumi.Input[str] fosid: FortiGate serial number.
        :param pulumi.Input[str] hostname: FortiGate hostname.
        :param pulumi.Input[str] name: FortiGate entry name.
        :param pulumi.Input[str] profile: FortiGate profile configuration.
        :param pulumi.Input[int] vdom: VDOM.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authorized is not None:
            pulumi.set(__self__, "authorized", authorized)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authorized(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[int]]:
        """
        device-id
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate serial number.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate hostname.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate profile configuration.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[int]]:
        """
        VDOM.
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vdom", value)

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


class Fortigate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        FortiGate controller configuration. Applies to FortiOS Version `>= 7.2.1`.

        ## Import

        ExtensionController Fortigate can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized: Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] device_id: device-id
        :param pulumi.Input[str] fosid: FortiGate serial number.
        :param pulumi.Input[str] hostname: FortiGate hostname.
        :param pulumi.Input[str] name: FortiGate entry name.
        :param pulumi.Input[str] profile: FortiGate profile configuration.
        :param pulumi.Input[int] vdom: VDOM.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortigateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        FortiGate controller configuration. Applies to FortiOS Version `>= 7.2.1`.

        ## Import

        ExtensionController Fortigate can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigate:Fortigate labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FortigateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortigateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortigateArgs.__new__(FortigateArgs)

            __props__.__dict__["authorized"] = authorized
            __props__.__dict__["description"] = description
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["name"] = name
            __props__.__dict__["profile"] = profile
            __props__.__dict__["vdom"] = vdom
            __props__.__dict__["vdomparam"] = vdomparam
        super(Fortigate, __self__).__init__(
            'fortios:extensioncontroller/fortigate:Fortigate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            device_id: Optional[pulumi.Input[int]] = None,
            fosid: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            profile: Optional[pulumi.Input[str]] = None,
            vdom: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Fortigate':
        """
        Get an existing Fortigate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized: Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[int] device_id: device-id
        :param pulumi.Input[str] fosid: FortiGate serial number.
        :param pulumi.Input[str] hostname: FortiGate hostname.
        :param pulumi.Input[str] name: FortiGate entry name.
        :param pulumi.Input[str] profile: FortiGate profile configuration.
        :param pulumi.Input[int] vdom: VDOM.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortigateState.__new__(_FortigateState)

        __props__.__dict__["authorized"] = authorized
        __props__.__dict__["description"] = description
        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["name"] = name
        __props__.__dict__["profile"] = profile
        __props__.__dict__["vdom"] = vdom
        __props__.__dict__["vdomparam"] = vdomparam
        return Fortigate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "authorized")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[int]:
        """
        device-id
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[str]:
        """
        FortiGate serial number.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        FortiGate hostname.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FortiGate entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[str]:
        """
        FortiGate profile configuration.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[int]:
        """
        VDOM.
        """
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
