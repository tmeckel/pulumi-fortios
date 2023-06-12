# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'Ospf6interfaceIpsecKey',
    'Ospf6interfaceNeighbor',
]

@pulumi.output_type
class Ospf6interfaceIpsecKey(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "authKey":
            suggest = "auth_key"
        elif key == "encKey":
            suggest = "enc_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Ospf6interfaceIpsecKey. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Ospf6interfaceIpsecKey.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Ospf6interfaceIpsecKey.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auth_key: Optional[str] = None,
                 enc_key: Optional[str] = None,
                 spi: Optional[int] = None):
        """
        :param str auth_key: Authentication key.
        :param str enc_key: Encryption key.
        :param int spi: Security Parameters Index.
        """
        if auth_key is not None:
            pulumi.set(__self__, "auth_key", auth_key)
        if enc_key is not None:
            pulumi.set(__self__, "enc_key", enc_key)
        if spi is not None:
            pulumi.set(__self__, "spi", spi)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> Optional[str]:
        """
        Authentication key.
        """
        return pulumi.get(self, "auth_key")

    @property
    @pulumi.getter(name="encKey")
    def enc_key(self) -> Optional[str]:
        """
        Encryption key.
        """
        return pulumi.get(self, "enc_key")

    @property
    @pulumi.getter
    def spi(self) -> Optional[int]:
        """
        Security Parameters Index.
        """
        return pulumi.get(self, "spi")


@pulumi.output_type
class Ospf6interfaceNeighbor(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "pollInterval":
            suggest = "poll_interval"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Ospf6interfaceNeighbor. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Ospf6interfaceNeighbor.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Ospf6interfaceNeighbor.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cost: Optional[int] = None,
                 ip6: Optional[str] = None,
                 poll_interval: Optional[int] = None,
                 priority: Optional[int] = None):
        """
        :param int cost: Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        :param str ip6: IPv6 link local address of the neighbor.
        :param int poll_interval: Poll interval time in seconds.
        :param int priority: priority
        """
        if cost is not None:
            pulumi.set(__self__, "cost", cost)
        if ip6 is not None:
            pulumi.set(__self__, "ip6", ip6)
        if poll_interval is not None:
            pulumi.set(__self__, "poll_interval", poll_interval)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter
    def cost(self) -> Optional[int]:
        """
        Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        """
        return pulumi.get(self, "cost")

    @property
    @pulumi.getter
    def ip6(self) -> Optional[str]:
        """
        IPv6 link local address of the neighbor.
        """
        return pulumi.get(self, "ip6")

    @property
    @pulumi.getter(name="pollInterval")
    def poll_interval(self) -> Optional[int]:
        """
        Poll interval time in seconds.
        """
        return pulumi.get(self, "poll_interval")

    @property
    @pulumi.getter
    def priority(self) -> Optional[int]:
        """
        priority
        """
        return pulumi.get(self, "priority")

