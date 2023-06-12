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
    'ServerIpRange',
    'ServerPrefixRange',
]

@pulumi.output_type
class ServerIpRange(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endIp":
            suggest = "end_ip"
        elif key == "startIp":
            suggest = "start_ip"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServerIpRange. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServerIpRange.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServerIpRange.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 end_ip: Optional[str] = None,
                 id: Optional[int] = None,
                 start_ip: Optional[str] = None):
        """
        :param str end_ip: End of IP range.
        :param int id: ID.
        :param str start_ip: Start of IP range.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[str]:
        """
        End of IP range.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[str]:
        """
        Start of IP range.
        """
        return pulumi.get(self, "start_ip")


@pulumi.output_type
class ServerPrefixRange(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endPrefix":
            suggest = "end_prefix"
        elif key == "prefixLength":
            suggest = "prefix_length"
        elif key == "startPrefix":
            suggest = "start_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServerPrefixRange. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServerPrefixRange.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServerPrefixRange.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 end_prefix: Optional[str] = None,
                 id: Optional[int] = None,
                 prefix_length: Optional[int] = None,
                 start_prefix: Optional[str] = None):
        """
        :param str end_prefix: End of prefix range.
        :param int id: ID.
        :param int prefix_length: Prefix length.
        :param str start_prefix: Start of prefix range.
        """
        if end_prefix is not None:
            pulumi.set(__self__, "end_prefix", end_prefix)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if start_prefix is not None:
            pulumi.set(__self__, "start_prefix", start_prefix)

    @property
    @pulumi.getter(name="endPrefix")
    def end_prefix(self) -> Optional[str]:
        """
        End of prefix range.
        """
        return pulumi.get(self, "end_prefix")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[int]:
        """
        Prefix length.
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter(name="startPrefix")
    def start_prefix(self) -> Optional[str]:
        """
        Start of prefix range.
        """
        return pulumi.get(self, "start_prefix")

