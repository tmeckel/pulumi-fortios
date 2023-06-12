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
    'DomainfilterEntry',
    'ProfileDnsTranslation',
    'ProfileDomainFilter',
    'ProfileExternalIpBlocklist',
    'ProfileFtgdDns',
    'ProfileFtgdDnsFilter',
]

@pulumi.output_type
class DomainfilterEntry(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 domain: Optional[str] = None,
                 id: Optional[int] = None,
                 status: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param str action: Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
        :param str domain: Domain entries to be filtered.
        :param int id: Id.
        :param str status: Enable/disable this domain filter. Valid values: `enable`, `disable`.
        :param str type: DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        Domain entries to be filtered.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable this domain filter. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ProfileDnsTranslation(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addrType":
            suggest = "addr_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileDnsTranslation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileDnsTranslation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileDnsTranslation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 addr_type: Optional[str] = None,
                 dst: Optional[str] = None,
                 dst6: Optional[str] = None,
                 id: Optional[int] = None,
                 netmask: Optional[str] = None,
                 prefix: Optional[int] = None,
                 src: Optional[str] = None,
                 src6: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str addr_type: DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        :param str dst: IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
        :param str dst6: IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
        :param int id: ID.
        :param str netmask: If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
        :param int prefix: If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
        :param str src: IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
        :param str src6: IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
        :param str status: Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
        """
        if addr_type is not None:
            pulumi.set(__self__, "addr_type", addr_type)
        if dst is not None:
            pulumi.set(__self__, "dst", dst)
        if dst6 is not None:
            pulumi.set(__self__, "dst6", dst6)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if src is not None:
            pulumi.set(__self__, "src", src)
        if src6 is not None:
            pulumi.set(__self__, "src6", src6)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> Optional[str]:
        """
        DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        """
        return pulumi.get(self, "addr_type")

    @property
    @pulumi.getter
    def dst(self) -> Optional[str]:
        """
        IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
        """
        return pulumi.get(self, "dst")

    @property
    @pulumi.getter
    def dst6(self) -> Optional[str]:
        """
        IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
        """
        return pulumi.get(self, "dst6")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def netmask(self) -> Optional[str]:
        """
        If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def prefix(self) -> Optional[int]:
        """
        If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def src(self) -> Optional[str]:
        """
        IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
        """
        return pulumi.get(self, "src")

    @property
    @pulumi.getter
    def src6(self) -> Optional[str]:
        """
        IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
        """
        return pulumi.get(self, "src6")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ProfileDomainFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "domainFilterTable":
            suggest = "domain_filter_table"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileDomainFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileDomainFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileDomainFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 domain_filter_table: Optional[int] = None):
        """
        :param int domain_filter_table: DNS domain filter table ID.
        """
        if domain_filter_table is not None:
            pulumi.set(__self__, "domain_filter_table", domain_filter_table)

    @property
    @pulumi.getter(name="domainFilterTable")
    def domain_filter_table(self) -> Optional[int]:
        """
        DNS domain filter table ID.
        """
        return pulumi.get(self, "domain_filter_table")


@pulumi.output_type
class ProfileExternalIpBlocklist(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: External domain block list name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        External domain block list name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ProfileFtgdDns(dict):
    def __init__(__self__, *,
                 filters: Optional[Sequence['outputs.ProfileFtgdDnsFilter']] = None,
                 options: Optional[str] = None):
        """
        :param Sequence['ProfileFtgdDnsFilterArgs'] filters: FortiGuard DNS domain filters. The structure of `filters` block is documented below.
        :param str options: FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
        """
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if options is not None:
            pulumi.set(__self__, "options", options)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.ProfileFtgdDnsFilter']]:
        """
        FortiGuard DNS domain filters. The structure of `filters` block is documented below.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def options(self) -> Optional[str]:
        """
        FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
        """
        return pulumi.get(self, "options")


@pulumi.output_type
class ProfileFtgdDnsFilter(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 category: Optional[int] = None,
                 id: Optional[int] = None,
                 log: Optional[str] = None):
        """
        :param str action: Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
        :param int category: Category number.
        :param int id: ID number.
        :param str log: Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if log is not None:
            pulumi.set(__self__, "log", log)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def category(self) -> Optional[int]:
        """
        Category number.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID number.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")


