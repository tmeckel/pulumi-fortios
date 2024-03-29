# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FortisandboxArgs', 'Fortisandbox']

@pulumi.input_type
class FortisandboxArgs:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 forticloud: Optional[pulumi.Input[str]] = None,
                 inline_scan: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fortisandbox resource.
        :param pulumi.Input[str] email: Notifier email address.
        :param pulumi.Input[str] enc_algorithm: Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        :param pulumi.Input[str] forticloud: Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] inline_scan: Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] server: IPv4 or IPv6 address of the remote FortiSandbox.
        :param pulumi.Input[str] source_ip: Source IP address for communications to FortiSandbox.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] status: Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if enc_algorithm is not None:
            pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if forticloud is not None:
            pulumi.set(__self__, "forticloud", forticloud)
        if inline_scan is not None:
            pulumi.set(__self__, "inline_scan", inline_scan)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Notifier email address.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        """
        return pulumi.get(self, "enc_algorithm")

    @enc_algorithm.setter
    def enc_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_algorithm", value)

    @property
    @pulumi.getter
    def forticloud(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "forticloud")

    @forticloud.setter
    def forticloud(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticloud", value)

    @property
    @pulumi.getter(name="inlineScan")
    def inline_scan(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "inline_scan")

    @inline_scan.setter
    def inline_scan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inline_scan", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        """
        Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        """
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 or IPv6 address of the remote FortiSandbox.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to FortiSandbox.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
class _FortisandboxState:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 forticloud: Optional[pulumi.Input[str]] = None,
                 inline_scan: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Fortisandbox resources.
        :param pulumi.Input[str] email: Notifier email address.
        :param pulumi.Input[str] enc_algorithm: Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        :param pulumi.Input[str] forticloud: Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] inline_scan: Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] server: IPv4 or IPv6 address of the remote FortiSandbox.
        :param pulumi.Input[str] source_ip: Source IP address for communications to FortiSandbox.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] status: Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if enc_algorithm is not None:
            pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if forticloud is not None:
            pulumi.set(__self__, "forticloud", forticloud)
        if inline_scan is not None:
            pulumi.set(__self__, "inline_scan", inline_scan)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Notifier email address.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        """
        return pulumi.get(self, "enc_algorithm")

    @enc_algorithm.setter
    def enc_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_algorithm", value)

    @property
    @pulumi.getter
    def forticloud(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "forticloud")

    @forticloud.setter
    def forticloud(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticloud", value)

    @property
    @pulumi.getter(name="inlineScan")
    def inline_scan(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "inline_scan")

    @inline_scan.setter
    def inline_scan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inline_scan", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        """
        Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        """
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 or IPv6 address of the remote FortiSandbox.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to FortiSandbox.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class Fortisandbox(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 forticloud: Optional[pulumi.Input[str]] = None,
                 inline_scan: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSandbox.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Fortisandbox("trname",
            enc_algorithm="default",
            ssl_min_proto_version="default",
            status="disable")
        ```

        ## Import

        System Fortisandbox can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/fortisandbox:Fortisandbox labelname SystemFortisandbox
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/fortisandbox:Fortisandbox labelname SystemFortisandbox
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: Notifier email address.
        :param pulumi.Input[str] enc_algorithm: Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        :param pulumi.Input[str] forticloud: Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] inline_scan: Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] server: IPv4 or IPv6 address of the remote FortiSandbox.
        :param pulumi.Input[str] source_ip: Source IP address for communications to FortiSandbox.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] status: Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortisandboxArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSandbox.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Fortisandbox("trname",
            enc_algorithm="default",
            ssl_min_proto_version="default",
            status="disable")
        ```

        ## Import

        System Fortisandbox can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/fortisandbox:Fortisandbox labelname SystemFortisandbox
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/fortisandbox:Fortisandbox labelname SystemFortisandbox
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FortisandboxArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortisandboxArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 forticloud: Optional[pulumi.Input[str]] = None,
                 inline_scan: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortisandboxArgs.__new__(FortisandboxArgs)

            __props__.__dict__["email"] = email
            __props__.__dict__["enc_algorithm"] = enc_algorithm
            __props__.__dict__["forticloud"] = forticloud
            __props__.__dict__["inline_scan"] = inline_scan
            __props__.__dict__["interface"] = interface
            __props__.__dict__["interface_select_method"] = interface_select_method
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Fortisandbox, __self__).__init__(
            'fortios:sys/fortisandbox:Fortisandbox',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email: Optional[pulumi.Input[str]] = None,
            enc_algorithm: Optional[pulumi.Input[str]] = None,
            forticloud: Optional[pulumi.Input[str]] = None,
            inline_scan: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            interface_select_method: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Fortisandbox':
        """
        Get an existing Fortisandbox resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: Notifier email address.
        :param pulumi.Input[str] enc_algorithm: Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        :param pulumi.Input[str] forticloud: Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] inline_scan: Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] server: IPv4 or IPv6 address of the remote FortiSandbox.
        :param pulumi.Input[str] source_ip: Source IP address for communications to FortiSandbox.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] status: Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortisandboxState.__new__(_FortisandboxState)

        __props__.__dict__["email"] = email
        __props__.__dict__["enc_algorithm"] = enc_algorithm
        __props__.__dict__["forticloud"] = forticloud
        __props__.__dict__["inline_scan"] = inline_scan
        __props__.__dict__["interface"] = interface
        __props__.__dict__["interface_select_method"] = interface_select_method
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Fortisandbox(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        Notifier email address.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> pulumi.Output[str]:
        """
        Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
        """
        return pulumi.get(self, "enc_algorithm")

    @property
    @pulumi.getter
    def forticloud(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "forticloud")

    @property
    @pulumi.getter(name="inlineScan")
    def inline_scan(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "inline_scan")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> pulumi.Output[str]:
        """
        Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        """
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        IPv4 or IPv6 address of the remote FortiSandbox.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IP address for communications to FortiSandbox.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> pulumi.Output[str]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

