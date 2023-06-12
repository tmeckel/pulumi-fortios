# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DevicemanagerInstallDeviceArgs', 'DevicemanagerInstallDevice']

@pulumi.input_type
class DevicemanagerInstallDeviceArgs:
    def __init__(__self__, *,
                 target_devname: pulumi.Input[str],
                 adom: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdom: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DevicemanagerInstallDevice resource.
        :param pulumi.Input[str] target_devname: Target device name.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[int] timeout: Timeout for installing the script to the target, default: 3 minutes.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        pulumi.set(__self__, "target_devname", target_devname)
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)

    @property
    @pulumi.getter(name="targetDevname")
    def target_devname(self) -> pulumi.Input[str]:
        """
        Target device name.
        """
        return pulumi.get(self, "target_devname")

    @target_devname.setter
    def target_devname(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_devname", value)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        Source ADOM name. default is 'root'
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for installing the script to the target, default: 3 minutes.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Vdom of managed device. default is 'root'
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)


@pulumi.input_type
class _DevicemanagerInstallDeviceState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 target_devname: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdom: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DevicemanagerInstallDevice resources.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] target_devname: Target device name.
        :param pulumi.Input[int] timeout: Timeout for installing the script to the target, default: 3 minutes.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if target_devname is not None:
            pulumi.set(__self__, "target_devname", target_devname)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        Source ADOM name. default is 'root'
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="targetDevname")
    def target_devname(self) -> Optional[pulumi.Input[str]]:
        """
        Target device name.
        """
        return pulumi.get(self, "target_devname")

    @target_devname.setter
    def target_devname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_devname", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for installing the script to the target, default: 3 minutes.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Vdom of managed device. default is 'root'
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)


class DevicemanagerInstallDevice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 target_devname: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports installing devicemanager script from FortiManager to the related device

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.DevicemanagerInstallDevice("test1",
            target_devname="FGVM64-test",
            timeout=5)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] target_devname: Target device name.
        :param pulumi.Input[int] timeout: Timeout for installing the script to the target, default: 3 minutes.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DevicemanagerInstallDeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports installing devicemanager script from FortiManager to the related device

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.DevicemanagerInstallDevice("test1",
            target_devname="FGVM64-test",
            timeout=5)
        ```

        :param str resource_name: The name of the resource.
        :param DevicemanagerInstallDeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DevicemanagerInstallDeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 target_devname: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DevicemanagerInstallDeviceArgs.__new__(DevicemanagerInstallDeviceArgs)

            __props__.__dict__["adom"] = adom
            if target_devname is None and not opts.urn:
                raise TypeError("Missing required property 'target_devname'")
            __props__.__dict__["target_devname"] = target_devname
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["vdom"] = vdom
        super(DevicemanagerInstallDevice, __self__).__init__(
            'fortios:fmg/devicemanagerInstallDevice:DevicemanagerInstallDevice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            target_devname: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            vdom: Optional[pulumi.Input[str]] = None) -> 'DevicemanagerInstallDevice':
        """
        Get an existing DevicemanagerInstallDevice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] target_devname: Target device name.
        :param pulumi.Input[int] timeout: Timeout for installing the script to the target, default: 3 minutes.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DevicemanagerInstallDeviceState.__new__(_DevicemanagerInstallDeviceState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["target_devname"] = target_devname
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["vdom"] = vdom
        return DevicemanagerInstallDevice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        """
        Source ADOM name. default is 'root'
        """
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="targetDevname")
    def target_devname(self) -> pulumi.Output[str]:
        """
        Target device name.
        """
        return pulumi.get(self, "target_devname")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Timeout for installing the script to the target, default: 3 minutes.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[Optional[str]]:
        """
        Vdom of managed device. default is 'root'
        """
        return pulumi.get(self, "vdom")
