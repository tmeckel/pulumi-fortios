# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FssopollingArgs', 'Fssopolling']

@pulumi.input_type
class FssopollingArgs:
    def __init__(__self__, *,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 listening_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fssopolling resource.
        :param pulumi.Input[str] auth_password: Password to connect to FSSO Agent.
        :param pulumi.Input[str] authentication: Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] listening_port: Listening port to accept clients (1 - 65535).
        :param pulumi.Input[str] status: Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_password is not None:
            pulumi.set(__self__, "auth_password", auth_password)
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if listening_port is not None:
            pulumi.set(__self__, "listening_port", listening_port)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> Optional[pulumi.Input[str]]:
        """
        Password to connect to FSSO Agent.
        """
        return pulumi.get(self, "auth_password")

    @auth_password.setter
    def auth_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_password", value)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter(name="listeningPort")
    def listening_port(self) -> Optional[pulumi.Input[int]]:
        """
        Listening port to accept clients (1 - 65535).
        """
        return pulumi.get(self, "listening_port")

    @listening_port.setter
    def listening_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listening_port", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
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
class _FssopollingState:
    def __init__(__self__, *,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 listening_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Fssopolling resources.
        :param pulumi.Input[str] auth_password: Password to connect to FSSO Agent.
        :param pulumi.Input[str] authentication: Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] listening_port: Listening port to accept clients (1 - 65535).
        :param pulumi.Input[str] status: Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_password is not None:
            pulumi.set(__self__, "auth_password", auth_password)
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if listening_port is not None:
            pulumi.set(__self__, "listening_port", listening_port)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> Optional[pulumi.Input[str]]:
        """
        Password to connect to FSSO Agent.
        """
        return pulumi.get(self, "auth_password")

    @auth_password.setter
    def auth_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_password", value)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter(name="listeningPort")
    def listening_port(self) -> Optional[pulumi.Input[int]]:
        """
        Listening port to accept clients (1 - 65535).
        """
        return pulumi.get(self, "listening_port")

    @listening_port.setter
    def listening_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listening_port", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
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


class Fssopolling(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 listening_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Fortinet Single Sign On (FSSO) server.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Fssopolling("trname",
            authentication="disable",
            listening_port=8000,
            status="enable")
        ```

        ## Import

        System FssoPolling can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/fssopolling:Fssopolling labelname SystemFssoPolling
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/fssopolling:Fssopolling labelname SystemFssoPolling
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_password: Password to connect to FSSO Agent.
        :param pulumi.Input[str] authentication: Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] listening_port: Listening port to accept clients (1 - 65535).
        :param pulumi.Input[str] status: Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FssopollingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Fortinet Single Sign On (FSSO) server.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Fssopolling("trname",
            authentication="disable",
            listening_port=8000,
            status="enable")
        ```

        ## Import

        System FssoPolling can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/fssopolling:Fssopolling labelname SystemFssoPolling
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/fssopolling:Fssopolling labelname SystemFssoPolling
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FssopollingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FssopollingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_password: Optional[pulumi.Input[str]] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 listening_port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FssopollingArgs.__new__(FssopollingArgs)

            __props__.__dict__["auth_password"] = None if auth_password is None else pulumi.Output.secret(auth_password)
            __props__.__dict__["authentication"] = authentication
            __props__.__dict__["listening_port"] = listening_port
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["authPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Fssopolling, __self__).__init__(
            'fortios:sys/fssopolling:Fssopolling',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_password: Optional[pulumi.Input[str]] = None,
            authentication: Optional[pulumi.Input[str]] = None,
            listening_port: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Fssopolling':
        """
        Get an existing Fssopolling resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_password: Password to connect to FSSO Agent.
        :param pulumi.Input[str] authentication: Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] listening_port: Listening port to accept clients (1 - 65535).
        :param pulumi.Input[str] status: Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FssopollingState.__new__(_FssopollingState)

        __props__.__dict__["auth_password"] = auth_password
        __props__.__dict__["authentication"] = authentication
        __props__.__dict__["listening_port"] = listening_port
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Fssopolling(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> pulumi.Output[Optional[str]]:
        """
        Password to connect to FSSO Agent.
        """
        return pulumi.get(self, "auth_password")

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output[str]:
        """
        Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="listeningPort")
    def listening_port(self) -> pulumi.Output[int]:
        """
        Listening port to accept clients (1 - 65535).
        """
        return pulumi.get(self, "listening_port")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

