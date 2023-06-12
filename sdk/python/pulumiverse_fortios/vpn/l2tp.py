# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['L2tpArgs', 'L2tp']

@pulumi.input_type
class L2tpArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 compress: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 enforce_ipsec: Optional[pulumi.Input[str]] = None,
                 hello_interval: Optional[pulumi.Input[int]] = None,
                 lcp_echo_interval: Optional[pulumi.Input[int]] = None,
                 lcp_max_echo_fails: Optional[pulumi.Input[int]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a L2tp resource.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] compress: Enable/disable data compression. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] enforce_ipsec: Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hello_interval: L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        :param pulumi.Input[int] lcp_echo_interval: Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        :param pulumi.Input[int] lcp_max_echo_fails: Maximum number of missed LCP echo messages before disconnect.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "status", status)
        if compress is not None:
            pulumi.set(__self__, "compress", compress)
        if eip is not None:
            pulumi.set(__self__, "eip", eip)
        if enforce_ipsec is not None:
            pulumi.set(__self__, "enforce_ipsec", enforce_ipsec)
        if hello_interval is not None:
            pulumi.set(__self__, "hello_interval", hello_interval)
        if lcp_echo_interval is not None:
            pulumi.set(__self__, "lcp_echo_interval", lcp_echo_interval)
        if lcp_max_echo_fails is not None:
            pulumi.set(__self__, "lcp_max_echo_fails", lcp_max_echo_fails)
        if sip is not None:
            pulumi.set(__self__, "sip", sip)
        if usrgrp is not None:
            pulumi.set(__self__, "usrgrp", usrgrp)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def compress(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable data compression. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "compress")

    @compress.setter
    def compress(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compress", value)

    @property
    @pulumi.getter
    def eip(self) -> Optional[pulumi.Input[str]]:
        """
        End IP.
        """
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="enforceIpsec")
    def enforce_ipsec(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "enforce_ipsec")

    @enforce_ipsec.setter
    def enforce_ipsec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enforce_ipsec", value)

    @property
    @pulumi.getter(name="helloInterval")
    def hello_interval(self) -> Optional[pulumi.Input[int]]:
        """
        L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        """
        return pulumi.get(self, "hello_interval")

    @hello_interval.setter
    def hello_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hello_interval", value)

    @property
    @pulumi.getter(name="lcpEchoInterval")
    def lcp_echo_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        """
        return pulumi.get(self, "lcp_echo_interval")

    @lcp_echo_interval.setter
    def lcp_echo_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lcp_echo_interval", value)

    @property
    @pulumi.getter(name="lcpMaxEchoFails")
    def lcp_max_echo_fails(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of missed LCP echo messages before disconnect.
        """
        return pulumi.get(self, "lcp_max_echo_fails")

    @lcp_max_echo_fails.setter
    def lcp_max_echo_fails(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lcp_max_echo_fails", value)

    @property
    @pulumi.getter
    def sip(self) -> Optional[pulumi.Input[str]]:
        """
        Start IP.
        """
        return pulumi.get(self, "sip")

    @sip.setter
    def sip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip", value)

    @property
    @pulumi.getter
    def usrgrp(self) -> Optional[pulumi.Input[str]]:
        """
        User group.
        """
        return pulumi.get(self, "usrgrp")

    @usrgrp.setter
    def usrgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usrgrp", value)

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
class _L2tpState:
    def __init__(__self__, *,
                 compress: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 enforce_ipsec: Optional[pulumi.Input[str]] = None,
                 hello_interval: Optional[pulumi.Input[int]] = None,
                 lcp_echo_interval: Optional[pulumi.Input[int]] = None,
                 lcp_max_echo_fails: Optional[pulumi.Input[int]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering L2tp resources.
        :param pulumi.Input[str] compress: Enable/disable data compression. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] enforce_ipsec: Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hello_interval: L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        :param pulumi.Input[int] lcp_echo_interval: Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        :param pulumi.Input[int] lcp_max_echo_fails: Maximum number of missed LCP echo messages before disconnect.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if compress is not None:
            pulumi.set(__self__, "compress", compress)
        if eip is not None:
            pulumi.set(__self__, "eip", eip)
        if enforce_ipsec is not None:
            pulumi.set(__self__, "enforce_ipsec", enforce_ipsec)
        if hello_interval is not None:
            pulumi.set(__self__, "hello_interval", hello_interval)
        if lcp_echo_interval is not None:
            pulumi.set(__self__, "lcp_echo_interval", lcp_echo_interval)
        if lcp_max_echo_fails is not None:
            pulumi.set(__self__, "lcp_max_echo_fails", lcp_max_echo_fails)
        if sip is not None:
            pulumi.set(__self__, "sip", sip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if usrgrp is not None:
            pulumi.set(__self__, "usrgrp", usrgrp)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def compress(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable data compression. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "compress")

    @compress.setter
    def compress(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compress", value)

    @property
    @pulumi.getter
    def eip(self) -> Optional[pulumi.Input[str]]:
        """
        End IP.
        """
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="enforceIpsec")
    def enforce_ipsec(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "enforce_ipsec")

    @enforce_ipsec.setter
    def enforce_ipsec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enforce_ipsec", value)

    @property
    @pulumi.getter(name="helloInterval")
    def hello_interval(self) -> Optional[pulumi.Input[int]]:
        """
        L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        """
        return pulumi.get(self, "hello_interval")

    @hello_interval.setter
    def hello_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hello_interval", value)

    @property
    @pulumi.getter(name="lcpEchoInterval")
    def lcp_echo_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        """
        return pulumi.get(self, "lcp_echo_interval")

    @lcp_echo_interval.setter
    def lcp_echo_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lcp_echo_interval", value)

    @property
    @pulumi.getter(name="lcpMaxEchoFails")
    def lcp_max_echo_fails(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of missed LCP echo messages before disconnect.
        """
        return pulumi.get(self, "lcp_max_echo_fails")

    @lcp_max_echo_fails.setter
    def lcp_max_echo_fails(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lcp_max_echo_fails", value)

    @property
    @pulumi.getter
    def sip(self) -> Optional[pulumi.Input[str]]:
        """
        Start IP.
        """
        return pulumi.get(self, "sip")

    @sip.setter
    def sip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def usrgrp(self) -> Optional[pulumi.Input[str]]:
        """
        User group.
        """
        return pulumi.get(self, "usrgrp")

    @usrgrp.setter
    def usrgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usrgrp", value)

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


class L2tp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compress: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 enforce_ipsec: Optional[pulumi.Input[str]] = None,
                 hello_interval: Optional[pulumi.Input[int]] = None,
                 lcp_echo_interval: Optional[pulumi.Input[int]] = None,
                 lcp_max_echo_fails: Optional[pulumi.Input[int]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure L2TP.

        ## Import

        Vpn L2Tp can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:vpn/l2tp:L2tp labelname VpnL2Tp
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:vpn/l2tp:L2tp labelname VpnL2Tp
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compress: Enable/disable data compression. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] enforce_ipsec: Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hello_interval: L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        :param pulumi.Input[int] lcp_echo_interval: Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        :param pulumi.Input[int] lcp_max_echo_fails: Maximum number of missed LCP echo messages before disconnect.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: L2tpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure L2TP.

        ## Import

        Vpn L2Tp can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:vpn/l2tp:L2tp labelname VpnL2Tp
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:vpn/l2tp:L2tp labelname VpnL2Tp
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param L2tpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(L2tpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compress: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 enforce_ipsec: Optional[pulumi.Input[str]] = None,
                 hello_interval: Optional[pulumi.Input[int]] = None,
                 lcp_echo_interval: Optional[pulumi.Input[int]] = None,
                 lcp_max_echo_fails: Optional[pulumi.Input[int]] = None,
                 sip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usrgrp: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = L2tpArgs.__new__(L2tpArgs)

            __props__.__dict__["compress"] = compress
            __props__.__dict__["eip"] = eip
            __props__.__dict__["enforce_ipsec"] = enforce_ipsec
            __props__.__dict__["hello_interval"] = hello_interval
            __props__.__dict__["lcp_echo_interval"] = lcp_echo_interval
            __props__.__dict__["lcp_max_echo_fails"] = lcp_max_echo_fails
            __props__.__dict__["sip"] = sip
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["usrgrp"] = usrgrp
            __props__.__dict__["vdomparam"] = vdomparam
        super(L2tp, __self__).__init__(
            'fortios:vpn/l2tp:L2tp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compress: Optional[pulumi.Input[str]] = None,
            eip: Optional[pulumi.Input[str]] = None,
            enforce_ipsec: Optional[pulumi.Input[str]] = None,
            hello_interval: Optional[pulumi.Input[int]] = None,
            lcp_echo_interval: Optional[pulumi.Input[int]] = None,
            lcp_max_echo_fails: Optional[pulumi.Input[int]] = None,
            sip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            usrgrp: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'L2tp':
        """
        Get an existing L2tp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compress: Enable/disable data compression. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] eip: End IP.
        :param pulumi.Input[str] enforce_ipsec: Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hello_interval: L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        :param pulumi.Input[int] lcp_echo_interval: Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        :param pulumi.Input[int] lcp_max_echo_fails: Maximum number of missed LCP echo messages before disconnect.
        :param pulumi.Input[str] sip: Start IP.
        :param pulumi.Input[str] status: Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usrgrp: User group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _L2tpState.__new__(_L2tpState)

        __props__.__dict__["compress"] = compress
        __props__.__dict__["eip"] = eip
        __props__.__dict__["enforce_ipsec"] = enforce_ipsec
        __props__.__dict__["hello_interval"] = hello_interval
        __props__.__dict__["lcp_echo_interval"] = lcp_echo_interval
        __props__.__dict__["lcp_max_echo_fails"] = lcp_max_echo_fails
        __props__.__dict__["sip"] = sip
        __props__.__dict__["status"] = status
        __props__.__dict__["usrgrp"] = usrgrp
        __props__.__dict__["vdomparam"] = vdomparam
        return L2tp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def compress(self) -> pulumi.Output[str]:
        """
        Enable/disable data compression. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "compress")

    @property
    @pulumi.getter
    def eip(self) -> pulumi.Output[str]:
        """
        End IP.
        """
        return pulumi.get(self, "eip")

    @property
    @pulumi.getter(name="enforceIpsec")
    def enforce_ipsec(self) -> pulumi.Output[str]:
        """
        Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "enforce_ipsec")

    @property
    @pulumi.getter(name="helloInterval")
    def hello_interval(self) -> pulumi.Output[int]:
        """
        L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
        """
        return pulumi.get(self, "hello_interval")

    @property
    @pulumi.getter(name="lcpEchoInterval")
    def lcp_echo_interval(self) -> pulumi.Output[int]:
        """
        Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
        """
        return pulumi.get(self, "lcp_echo_interval")

    @property
    @pulumi.getter(name="lcpMaxEchoFails")
    def lcp_max_echo_fails(self) -> pulumi.Output[int]:
        """
        Maximum number of missed LCP echo messages before disconnect.
        """
        return pulumi.get(self, "lcp_max_echo_fails")

    @property
    @pulumi.getter
    def sip(self) -> pulumi.Output[str]:
        """
        Start IP.
        """
        return pulumi.get(self, "sip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def usrgrp(self) -> pulumi.Output[str]:
        """
        User group.
        """
        return pulumi.get(self, "usrgrp")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
