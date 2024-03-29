# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthgroupArgs', 'Authgroup']

@pulumi.input_type
class AuthgroupArgs:
    def __init__(__self__, *,
                 cert: pulumi.Input[str],
                 auth_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 peer_accept: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Authgroup resource.
        :param pulumi.Input[str] cert: Name of certificate to identify this peer.
        :param pulumi.Input[str] auth_method: Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        :param pulumi.Input[str] name: Auth-group name.
        :param pulumi.Input[str] peer: If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        :param pulumi.Input[str] peer_accept: Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        :param pulumi.Input[str] psk: Pre-shared key used by the peers in this authentication group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "cert", cert)
        if auth_method is not None:
            pulumi.set(__self__, "auth_method", auth_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer is not None:
            pulumi.set(__self__, "peer", peer)
        if peer_accept is not None:
            pulumi.set(__self__, "peer_accept", peer_accept)
        if psk is not None:
            pulumi.set(__self__, "psk", psk)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Input[str]:
        """
        Name of certificate to identify this peer.
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> Optional[pulumi.Input[str]]:
        """
        Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        """
        return pulumi.get(self, "auth_method")

    @auth_method.setter
    def auth_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Auth-group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def peer(self) -> Optional[pulumi.Input[str]]:
        """
        If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        """
        return pulumi.get(self, "peer")

    @peer.setter
    def peer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer", value)

    @property
    @pulumi.getter(name="peerAccept")
    def peer_accept(self) -> Optional[pulumi.Input[str]]:
        """
        Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        """
        return pulumi.get(self, "peer_accept")

    @peer_accept.setter
    def peer_accept(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_accept", value)

    @property
    @pulumi.getter
    def psk(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared key used by the peers in this authentication group.
        """
        return pulumi.get(self, "psk")

    @psk.setter
    def psk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psk", value)

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
class _AuthgroupState:
    def __init__(__self__, *,
                 auth_method: Optional[pulumi.Input[str]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 peer_accept: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Authgroup resources.
        :param pulumi.Input[str] auth_method: Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        :param pulumi.Input[str] cert: Name of certificate to identify this peer.
        :param pulumi.Input[str] name: Auth-group name.
        :param pulumi.Input[str] peer: If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        :param pulumi.Input[str] peer_accept: Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        :param pulumi.Input[str] psk: Pre-shared key used by the peers in this authentication group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth_method is not None:
            pulumi.set(__self__, "auth_method", auth_method)
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer is not None:
            pulumi.set(__self__, "peer", peer)
        if peer_accept is not None:
            pulumi.set(__self__, "peer_accept", peer_accept)
        if psk is not None:
            pulumi.set(__self__, "psk", psk)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> Optional[pulumi.Input[str]]:
        """
        Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        """
        return pulumi.get(self, "auth_method")

    @auth_method.setter
    def auth_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_method", value)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[str]]:
        """
        Name of certificate to identify this peer.
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Auth-group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def peer(self) -> Optional[pulumi.Input[str]]:
        """
        If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        """
        return pulumi.get(self, "peer")

    @peer.setter
    def peer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer", value)

    @property
    @pulumi.getter(name="peerAccept")
    def peer_accept(self) -> Optional[pulumi.Input[str]]:
        """
        Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        """
        return pulumi.get(self, "peer_accept")

    @peer_accept.setter
    def peer_accept(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_accept", value)

    @property
    @pulumi.getter
    def psk(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared key used by the peers in this authentication group.
        """
        return pulumi.get(self, "psk")

    @psk.setter
    def psk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psk", value)

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


class Authgroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method: Optional[pulumi.Input[str]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 peer_accept: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure WAN optimization authentication groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wanopt.Authgroup("trname",
            auth_method="cert",
            cert="Fortinet_CA_SSL",
            peer_accept="any")
        ```

        ## Import

        Wanopt AuthGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:wanopt/authgroup:Authgroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:wanopt/authgroup:Authgroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method: Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        :param pulumi.Input[str] cert: Name of certificate to identify this peer.
        :param pulumi.Input[str] name: Auth-group name.
        :param pulumi.Input[str] peer: If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        :param pulumi.Input[str] peer_accept: Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        :param pulumi.Input[str] psk: Pre-shared key used by the peers in this authentication group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthgroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WAN optimization authentication groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.wanopt.Authgroup("trname",
            auth_method="cert",
            cert="Fortinet_CA_SSL",
            peer_accept="any")
        ```

        ## Import

        Wanopt AuthGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:wanopt/authgroup:Authgroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:wanopt/authgroup:Authgroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param AuthgroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthgroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method: Optional[pulumi.Input[str]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer: Optional[pulumi.Input[str]] = None,
                 peer_accept: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthgroupArgs.__new__(AuthgroupArgs)

            __props__.__dict__["auth_method"] = auth_method
            if cert is None and not opts.urn:
                raise TypeError("Missing required property 'cert'")
            __props__.__dict__["cert"] = cert
            __props__.__dict__["name"] = name
            __props__.__dict__["peer"] = peer
            __props__.__dict__["peer_accept"] = peer_accept
            __props__.__dict__["psk"] = None if psk is None else pulumi.Output.secret(psk)
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["psk"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Authgroup, __self__).__init__(
            'fortios:wanopt/authgroup:Authgroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_method: Optional[pulumi.Input[str]] = None,
            cert: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            peer: Optional[pulumi.Input[str]] = None,
            peer_accept: Optional[pulumi.Input[str]] = None,
            psk: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Authgroup':
        """
        Get an existing Authgroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_method: Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        :param pulumi.Input[str] cert: Name of certificate to identify this peer.
        :param pulumi.Input[str] name: Auth-group name.
        :param pulumi.Input[str] peer: If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        :param pulumi.Input[str] peer_accept: Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        :param pulumi.Input[str] psk: Pre-shared key used by the peers in this authentication group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthgroupState.__new__(_AuthgroupState)

        __props__.__dict__["auth_method"] = auth_method
        __props__.__dict__["cert"] = cert
        __props__.__dict__["name"] = name
        __props__.__dict__["peer"] = peer
        __props__.__dict__["peer_accept"] = peer_accept
        __props__.__dict__["psk"] = psk
        __props__.__dict__["vdomparam"] = vdomparam
        return Authgroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> pulumi.Output[str]:
        """
        Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
        """
        return pulumi.get(self, "auth_method")

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output[str]:
        """
        Name of certificate to identify this peer.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Auth-group name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def peer(self) -> pulumi.Output[str]:
        """
        If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
        """
        return pulumi.get(self, "peer")

    @property
    @pulumi.getter(name="peerAccept")
    def peer_accept(self) -> pulumi.Output[str]:
        """
        Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
        """
        return pulumi.get(self, "peer_accept")

    @property
    @pulumi.getter
    def psk(self) -> pulumi.Output[Optional[str]]:
        """
        Pre-shared key used by the peers in this authentication group.
        """
        return pulumi.get(self, "psk")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

