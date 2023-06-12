# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetMobiletunnelResult',
    'AwaitableGetMobiletunnelResult',
    'get_mobiletunnel',
    'get_mobiletunnel_output',
]

@pulumi.output_type
class GetMobiletunnelResult:
    """
    A collection of values returned by getMobiletunnel.
    """
    def __init__(__self__, hash_algorithm=None, home_address=None, home_agent=None, id=None, lifetime=None, n_mhae_key=None, n_mhae_key_type=None, n_mhae_spi=None, name=None, networks=None, reg_interval=None, reg_retry=None, renew_interval=None, roaming_interface=None, status=None, tunnel_mode=None, vdomparam=None):
        if hash_algorithm and not isinstance(hash_algorithm, str):
            raise TypeError("Expected argument 'hash_algorithm' to be a str")
        pulumi.set(__self__, "hash_algorithm", hash_algorithm)
        if home_address and not isinstance(home_address, str):
            raise TypeError("Expected argument 'home_address' to be a str")
        pulumi.set(__self__, "home_address", home_address)
        if home_agent and not isinstance(home_agent, str):
            raise TypeError("Expected argument 'home_agent' to be a str")
        pulumi.set(__self__, "home_agent", home_agent)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lifetime and not isinstance(lifetime, int):
            raise TypeError("Expected argument 'lifetime' to be a int")
        pulumi.set(__self__, "lifetime", lifetime)
        if n_mhae_key and not isinstance(n_mhae_key, str):
            raise TypeError("Expected argument 'n_mhae_key' to be a str")
        pulumi.set(__self__, "n_mhae_key", n_mhae_key)
        if n_mhae_key_type and not isinstance(n_mhae_key_type, str):
            raise TypeError("Expected argument 'n_mhae_key_type' to be a str")
        pulumi.set(__self__, "n_mhae_key_type", n_mhae_key_type)
        if n_mhae_spi and not isinstance(n_mhae_spi, int):
            raise TypeError("Expected argument 'n_mhae_spi' to be a int")
        pulumi.set(__self__, "n_mhae_spi", n_mhae_spi)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if reg_interval and not isinstance(reg_interval, int):
            raise TypeError("Expected argument 'reg_interval' to be a int")
        pulumi.set(__self__, "reg_interval", reg_interval)
        if reg_retry and not isinstance(reg_retry, int):
            raise TypeError("Expected argument 'reg_retry' to be a int")
        pulumi.set(__self__, "reg_retry", reg_retry)
        if renew_interval and not isinstance(renew_interval, int):
            raise TypeError("Expected argument 'renew_interval' to be a int")
        pulumi.set(__self__, "renew_interval", renew_interval)
        if roaming_interface and not isinstance(roaming_interface, str):
            raise TypeError("Expected argument 'roaming_interface' to be a str")
        pulumi.set(__self__, "roaming_interface", roaming_interface)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tunnel_mode and not isinstance(tunnel_mode, str):
            raise TypeError("Expected argument 'tunnel_mode' to be a str")
        pulumi.set(__self__, "tunnel_mode", tunnel_mode)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> str:
        """
        Hash Algorithm (Keyed MD5).
        """
        return pulumi.get(self, "hash_algorithm")

    @property
    @pulumi.getter(name="homeAddress")
    def home_address(self) -> str:
        """
        Home IP address (Format: xxx.xxx.xxx.xxx).
        """
        return pulumi.get(self, "home_address")

    @property
    @pulumi.getter(name="homeAgent")
    def home_agent(self) -> str:
        """
        IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
        """
        return pulumi.get(self, "home_agent")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lifetime(self) -> int:
        """
        NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
        """
        return pulumi.get(self, "lifetime")

    @property
    @pulumi.getter(name="nMhaeKey")
    def n_mhae_key(self) -> str:
        """
        NEMO authentication key.
        """
        return pulumi.get(self, "n_mhae_key")

    @property
    @pulumi.getter(name="nMhaeKeyType")
    def n_mhae_key_type(self) -> str:
        """
        NEMO authentication key type (ascii or base64).
        """
        return pulumi.get(self, "n_mhae_key_type")

    @property
    @pulumi.getter(name="nMhaeSpi")
    def n_mhae_spi(self) -> int:
        """
        NEMO authentication SPI (default: 256).
        """
        return pulumi.get(self, "n_mhae_spi")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetMobiletunnelNetworkResult']:
        """
        NEMO network configuration. The structure of `network` block is documented below.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="regInterval")
    def reg_interval(self) -> int:
        """
        NMMO HA registration interval (5 - 300, default = 5).
        """
        return pulumi.get(self, "reg_interval")

    @property
    @pulumi.getter(name="regRetry")
    def reg_retry(self) -> int:
        """
        Maximum number of NMMO HA registration retries (1 to 30, default = 3).
        """
        return pulumi.get(self, "reg_retry")

    @property
    @pulumi.getter(name="renewInterval")
    def renew_interval(self) -> int:
        """
        Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
        """
        return pulumi.get(self, "renew_interval")

    @property
    @pulumi.getter(name="roamingInterface")
    def roaming_interface(self) -> str:
        """
        Select the associated interface name from available options.
        """
        return pulumi.get(self, "roaming_interface")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this mobile tunnel.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> str:
        """
        NEMO tunnnel mode (GRE tunnel).
        """
        return pulumi.get(self, "tunnel_mode")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetMobiletunnelResult(GetMobiletunnelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMobiletunnelResult(
            hash_algorithm=self.hash_algorithm,
            home_address=self.home_address,
            home_agent=self.home_agent,
            id=self.id,
            lifetime=self.lifetime,
            n_mhae_key=self.n_mhae_key,
            n_mhae_key_type=self.n_mhae_key_type,
            n_mhae_spi=self.n_mhae_spi,
            name=self.name,
            networks=self.networks,
            reg_interval=self.reg_interval,
            reg_retry=self.reg_retry,
            renew_interval=self.renew_interval,
            roaming_interface=self.roaming_interface,
            status=self.status,
            tunnel_mode=self.tunnel_mode,
            vdomparam=self.vdomparam)


def get_mobiletunnel(name: Optional[str] = None,
                     vdomparam: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMobiletunnelResult:
    """
    Use this data source to get information on an fortios system mobiletunnel


    :param str name: Specify the name of the desired system mobiletunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:sys/getMobiletunnel:getMobiletunnel', __args__, opts=opts, typ=GetMobiletunnelResult).value

    return AwaitableGetMobiletunnelResult(
        hash_algorithm=__ret__.hash_algorithm,
        home_address=__ret__.home_address,
        home_agent=__ret__.home_agent,
        id=__ret__.id,
        lifetime=__ret__.lifetime,
        n_mhae_key=__ret__.n_mhae_key,
        n_mhae_key_type=__ret__.n_mhae_key_type,
        n_mhae_spi=__ret__.n_mhae_spi,
        name=__ret__.name,
        networks=__ret__.networks,
        reg_interval=__ret__.reg_interval,
        reg_retry=__ret__.reg_retry,
        renew_interval=__ret__.renew_interval,
        roaming_interface=__ret__.roaming_interface,
        status=__ret__.status,
        tunnel_mode=__ret__.tunnel_mode,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_mobiletunnel)
def get_mobiletunnel_output(name: Optional[pulumi.Input[str]] = None,
                            vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMobiletunnelResult]:
    """
    Use this data source to get information on an fortios system mobiletunnel


    :param str name: Specify the name of the desired system mobiletunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
