# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomaincontrollerArgs', 'Domaincontroller']

@pulumi.input_type
class DomaincontrollerArgs:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Domaincontroller resource.
        :param pulumi.Input[str] domain_name: Fully qualified domain name (FQDN).
        :param pulumi.Input[str] hostname: Hostname of the server to connect to.
        :param pulumi.Input[str] ip: IPv4 server address.
        :param pulumi.Input[str] ip6: IPv6 server address.
        :param pulumi.Input[str] password: Password for specified username.
        :param pulumi.Input[int] port: Port number of service. Port number 0 indicates automatic discovery.
        :param pulumi.Input[str] server_name: Name of the server to connect to.
        :param pulumi.Input[str] username: User name to sign in with. Must have proper permissions for service.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip6 is not None:
            pulumi.set(__self__, "ip6", ip6)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Fully qualified domain name (FQDN).
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the server to connect to.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 server address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 server address.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for specified username.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number of service. Port number 0 indicates automatic discovery.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the server to connect to.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        User name to sign in with. Must have proper permissions for service.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

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
class _DomaincontrollerState:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Domaincontroller resources.
        :param pulumi.Input[str] domain_name: Fully qualified domain name (FQDN).
        :param pulumi.Input[str] hostname: Hostname of the server to connect to.
        :param pulumi.Input[str] ip: IPv4 server address.
        :param pulumi.Input[str] ip6: IPv6 server address.
        :param pulumi.Input[str] password: Password for specified username.
        :param pulumi.Input[int] port: Port number of service. Port number 0 indicates automatic discovery.
        :param pulumi.Input[str] server_name: Name of the server to connect to.
        :param pulumi.Input[str] username: User name to sign in with. Must have proper permissions for service.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip6 is not None:
            pulumi.set(__self__, "ip6", ip6)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Fully qualified domain name (FQDN).
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the server to connect to.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 server address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 server address.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for specified username.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number of service. Port number 0 indicates automatic discovery.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the server to connect to.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        User name to sign in with. Must have proper permissions for service.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

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


class Domaincontroller(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Define known domain controller servers. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.

        ## Import

        CredentialStore DomainController can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:credentialstore/domaincontroller:Domaincontroller labelname {{server_name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:credentialstore/domaincontroller:Domaincontroller labelname {{server_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Fully qualified domain name (FQDN).
        :param pulumi.Input[str] hostname: Hostname of the server to connect to.
        :param pulumi.Input[str] ip: IPv4 server address.
        :param pulumi.Input[str] ip6: IPv6 server address.
        :param pulumi.Input[str] password: Password for specified username.
        :param pulumi.Input[int] port: Port number of service. Port number 0 indicates automatic discovery.
        :param pulumi.Input[str] server_name: Name of the server to connect to.
        :param pulumi.Input[str] username: User name to sign in with. Must have proper permissions for service.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DomaincontrollerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Define known domain controller servers. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.

        ## Import

        CredentialStore DomainController can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:credentialstore/domaincontroller:Domaincontroller labelname {{server_name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:credentialstore/domaincontroller:Domaincontroller labelname {{server_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param DomaincontrollerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomaincontrollerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomaincontrollerArgs.__new__(DomaincontrollerArgs)

            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["ip"] = ip
            __props__.__dict__["ip6"] = ip6
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["port"] = port
            __props__.__dict__["server_name"] = server_name
            __props__.__dict__["username"] = username
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Domaincontroller, __self__).__init__(
            'fortios:credentialstore/domaincontroller:Domaincontroller',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip6: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Domaincontroller':
        """
        Get an existing Domaincontroller resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Fully qualified domain name (FQDN).
        :param pulumi.Input[str] hostname: Hostname of the server to connect to.
        :param pulumi.Input[str] ip: IPv4 server address.
        :param pulumi.Input[str] ip6: IPv6 server address.
        :param pulumi.Input[str] password: Password for specified username.
        :param pulumi.Input[int] port: Port number of service. Port number 0 indicates automatic discovery.
        :param pulumi.Input[str] server_name: Name of the server to connect to.
        :param pulumi.Input[str] username: User name to sign in with. Must have proper permissions for service.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomaincontrollerState.__new__(_DomaincontrollerState)

        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip6"] = ip6
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["server_name"] = server_name
        __props__.__dict__["username"] = username
        __props__.__dict__["vdomparam"] = vdomparam
        return Domaincontroller(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Fully qualified domain name (FQDN).
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        Hostname of the server to connect to.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IPv4 server address.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def ip6(self) -> pulumi.Output[str]:
        """
        IPv6 server address.
        """
        return pulumi.get(self, "ip6")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password for specified username.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port number of service. Port number 0 indicates automatic discovery.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[str]:
        """
        Name of the server to connect to.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        User name to sign in with. Must have proper permissions for service.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

