# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SettingsArgs', 'Settings']

@pulumi.input_type
class SettingsArgs:
    def __init__(__self__, *,
                 ips_packet_quota: Optional[pulumi.Input[int]] = None,
                 packet_log_history: Optional[pulumi.Input[int]] = None,
                 packet_log_memory: Optional[pulumi.Input[int]] = None,
                 packet_log_post_attack: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Settings resource.
        :param pulumi.Input[int] ips_packet_quota: Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        :param pulumi.Input[int] packet_log_history: Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        :param pulumi.Input[int] packet_log_memory: Maximum memory can be used by packet log (64 - 8192 kB).
        :param pulumi.Input[int] packet_log_post_attack: Number of packets to log after the IPS signature is detected (0 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ips_packet_quota is not None:
            pulumi.set(__self__, "ips_packet_quota", ips_packet_quota)
        if packet_log_history is not None:
            pulumi.set(__self__, "packet_log_history", packet_log_history)
        if packet_log_memory is not None:
            pulumi.set(__self__, "packet_log_memory", packet_log_memory)
        if packet_log_post_attack is not None:
            pulumi.set(__self__, "packet_log_post_attack", packet_log_post_attack)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="ipsPacketQuota")
    def ips_packet_quota(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        """
        return pulumi.get(self, "ips_packet_quota")

    @ips_packet_quota.setter
    def ips_packet_quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ips_packet_quota", value)

    @property
    @pulumi.getter(name="packetLogHistory")
    def packet_log_history(self) -> Optional[pulumi.Input[int]]:
        """
        Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        """
        return pulumi.get(self, "packet_log_history")

    @packet_log_history.setter
    def packet_log_history(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_log_history", value)

    @property
    @pulumi.getter(name="packetLogMemory")
    def packet_log_memory(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum memory can be used by packet log (64 - 8192 kB).
        """
        return pulumi.get(self, "packet_log_memory")

    @packet_log_memory.setter
    def packet_log_memory(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_log_memory", value)

    @property
    @pulumi.getter(name="packetLogPostAttack")
    def packet_log_post_attack(self) -> Optional[pulumi.Input[int]]:
        """
        Number of packets to log after the IPS signature is detected (0 - 255).
        """
        return pulumi.get(self, "packet_log_post_attack")

    @packet_log_post_attack.setter
    def packet_log_post_attack(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_log_post_attack", value)

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
class _SettingsState:
    def __init__(__self__, *,
                 ips_packet_quota: Optional[pulumi.Input[int]] = None,
                 packet_log_history: Optional[pulumi.Input[int]] = None,
                 packet_log_memory: Optional[pulumi.Input[int]] = None,
                 packet_log_post_attack: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Settings resources.
        :param pulumi.Input[int] ips_packet_quota: Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        :param pulumi.Input[int] packet_log_history: Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        :param pulumi.Input[int] packet_log_memory: Maximum memory can be used by packet log (64 - 8192 kB).
        :param pulumi.Input[int] packet_log_post_attack: Number of packets to log after the IPS signature is detected (0 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ips_packet_quota is not None:
            pulumi.set(__self__, "ips_packet_quota", ips_packet_quota)
        if packet_log_history is not None:
            pulumi.set(__self__, "packet_log_history", packet_log_history)
        if packet_log_memory is not None:
            pulumi.set(__self__, "packet_log_memory", packet_log_memory)
        if packet_log_post_attack is not None:
            pulumi.set(__self__, "packet_log_post_attack", packet_log_post_attack)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="ipsPacketQuota")
    def ips_packet_quota(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        """
        return pulumi.get(self, "ips_packet_quota")

    @ips_packet_quota.setter
    def ips_packet_quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ips_packet_quota", value)

    @property
    @pulumi.getter(name="packetLogHistory")
    def packet_log_history(self) -> Optional[pulumi.Input[int]]:
        """
        Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        """
        return pulumi.get(self, "packet_log_history")

    @packet_log_history.setter
    def packet_log_history(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_log_history", value)

    @property
    @pulumi.getter(name="packetLogMemory")
    def packet_log_memory(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum memory can be used by packet log (64 - 8192 kB).
        """
        return pulumi.get(self, "packet_log_memory")

    @packet_log_memory.setter
    def packet_log_memory(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_log_memory", value)

    @property
    @pulumi.getter(name="packetLogPostAttack")
    def packet_log_post_attack(self) -> Optional[pulumi.Input[int]]:
        """
        Number of packets to log after the IPS signature is detected (0 - 255).
        """
        return pulumi.get(self, "packet_log_post_attack")

    @packet_log_post_attack.setter
    def packet_log_post_attack(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_log_post_attack", value)

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


class Settings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ips_packet_quota: Optional[pulumi.Input[int]] = None,
                 packet_log_history: Optional[pulumi.Input[int]] = None,
                 packet_log_memory: Optional[pulumi.Input[int]] = None,
                 packet_log_post_attack: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPS VDOM parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.ips.Settings("trname",
            ips_packet_quota=0,
            packet_log_history=1,
            packet_log_memory=256,
            packet_log_post_attack=0)
        ```

        ## Import

        Ips Settings can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:ips/settings:Settings labelname IpsSettings
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:ips/settings:Settings labelname IpsSettings
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ips_packet_quota: Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        :param pulumi.Input[int] packet_log_history: Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        :param pulumi.Input[int] packet_log_memory: Maximum memory can be used by packet log (64 - 8192 kB).
        :param pulumi.Input[int] packet_log_post_attack: Number of packets to log after the IPS signature is detected (0 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPS VDOM parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.ips.Settings("trname",
            ips_packet_quota=0,
            packet_log_history=1,
            packet_log_memory=256,
            packet_log_post_attack=0)
        ```

        ## Import

        Ips Settings can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:ips/settings:Settings labelname IpsSettings
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:ips/settings:Settings labelname IpsSettings
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ips_packet_quota: Optional[pulumi.Input[int]] = None,
                 packet_log_history: Optional[pulumi.Input[int]] = None,
                 packet_log_memory: Optional[pulumi.Input[int]] = None,
                 packet_log_post_attack: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SettingsArgs.__new__(SettingsArgs)

            __props__.__dict__["ips_packet_quota"] = ips_packet_quota
            __props__.__dict__["packet_log_history"] = packet_log_history
            __props__.__dict__["packet_log_memory"] = packet_log_memory
            __props__.__dict__["packet_log_post_attack"] = packet_log_post_attack
            __props__.__dict__["vdomparam"] = vdomparam
        super(Settings, __self__).__init__(
            'fortios:ips/settings:Settings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ips_packet_quota: Optional[pulumi.Input[int]] = None,
            packet_log_history: Optional[pulumi.Input[int]] = None,
            packet_log_memory: Optional[pulumi.Input[int]] = None,
            packet_log_post_attack: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Settings':
        """
        Get an existing Settings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] ips_packet_quota: Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        :param pulumi.Input[int] packet_log_history: Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        :param pulumi.Input[int] packet_log_memory: Maximum memory can be used by packet log (64 - 8192 kB).
        :param pulumi.Input[int] packet_log_post_attack: Number of packets to log after the IPS signature is detected (0 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SettingsState.__new__(_SettingsState)

        __props__.__dict__["ips_packet_quota"] = ips_packet_quota
        __props__.__dict__["packet_log_history"] = packet_log_history
        __props__.__dict__["packet_log_memory"] = packet_log_memory
        __props__.__dict__["packet_log_post_attack"] = packet_log_post_attack
        __props__.__dict__["vdomparam"] = vdomparam
        return Settings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipsPacketQuota")
    def ips_packet_quota(self) -> pulumi.Output[int]:
        """
        Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        """
        return pulumi.get(self, "ips_packet_quota")

    @property
    @pulumi.getter(name="packetLogHistory")
    def packet_log_history(self) -> pulumi.Output[int]:
        """
        Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        """
        return pulumi.get(self, "packet_log_history")

    @property
    @pulumi.getter(name="packetLogMemory")
    def packet_log_memory(self) -> pulumi.Output[int]:
        """
        Maximum memory can be used by packet log (64 - 8192 kB).
        """
        return pulumi.get(self, "packet_log_memory")

    @property
    @pulumi.getter(name="packetLogPostAttack")
    def packet_log_post_attack(self) -> pulumi.Output[int]:
        """
        Number of packets to log after the IPS signature is detected (0 - 255).
        """
        return pulumi.get(self, "packet_log_post_attack")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

