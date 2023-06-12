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
    'Ospf6interfaceIpsecKeyArgs',
    'Ospf6interfaceNeighborArgs',
]

@pulumi.input_type
class Ospf6interfaceIpsecKeyArgs:
    def __init__(__self__, *,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 enc_key: Optional[pulumi.Input[str]] = None,
                 spi: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] auth_key: Authentication key.
        :param pulumi.Input[str] enc_key: Encryption key.
        :param pulumi.Input[int] spi: Security Parameters Index.
        """
        if auth_key is not None:
            pulumi.set(__self__, "auth_key", auth_key)
        if enc_key is not None:
            pulumi.set(__self__, "enc_key", enc_key)
        if spi is not None:
            pulumi.set(__self__, "spi", spi)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication key.
        """
        return pulumi.get(self, "auth_key")

    @auth_key.setter
    def auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_key", value)

    @property
    @pulumi.getter(name="encKey")
    def enc_key(self) -> Optional[pulumi.Input[str]]:
        """
        Encryption key.
        """
        return pulumi.get(self, "enc_key")

    @enc_key.setter
    def enc_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_key", value)

    @property
    @pulumi.getter
    def spi(self) -> Optional[pulumi.Input[int]]:
        """
        Security Parameters Index.
        """
        return pulumi.get(self, "spi")

    @spi.setter
    def spi(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "spi", value)


@pulumi.input_type
class Ospf6interfaceNeighborArgs:
    def __init__(__self__, *,
                 cost: Optional[pulumi.Input[int]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 poll_interval: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] cost: Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        :param pulumi.Input[str] ip6: IPv6 link local address of the neighbor.
        :param pulumi.Input[int] poll_interval: Poll interval time in seconds.
        :param pulumi.Input[int] priority: priority
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
    def cost(self) -> Optional[pulumi.Input[int]]:
        """
        Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        """
        return pulumi.get(self, "cost")

    @cost.setter
    def cost(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cost", value)

    @property
    @pulumi.getter
    def ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 link local address of the neighbor.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter(name="pollInterval")
    def poll_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Poll interval time in seconds.
        """
        return pulumi.get(self, "poll_interval")

    @poll_interval.setter
    def poll_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "poll_interval", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        priority
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)


