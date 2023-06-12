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
    'DomainfilterEntryArgs',
    'ProfileDnsTranslationArgs',
    'ProfileDomainFilterArgs',
    'ProfileExternalIpBlocklistArgs',
    'ProfileFtgdDnsArgs',
    'ProfileFtgdDnsFilterArgs',
]

@pulumi.input_type
class DomainfilterEntryArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
        :param pulumi.Input[str] domain: Domain entries to be filtered.
        :param pulumi.Input[int] id: Id.
        :param pulumi.Input[str] status: Enable/disable this domain filter. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
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
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain entries to be filtered.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Id.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this domain filter. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProfileDnsTranslationArgs:
    def __init__(__self__, *,
                 addr_type: Optional[pulumi.Input[str]] = None,
                 dst: Optional[pulumi.Input[str]] = None,
                 dst6: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[int]] = None,
                 src: Optional[pulumi.Input[str]] = None,
                 src6: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] addr_type: DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        :param pulumi.Input[str] dst: IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
        :param pulumi.Input[str] dst6: IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] netmask: If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
        :param pulumi.Input[int] prefix: If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
        :param pulumi.Input[str] src: IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
        :param pulumi.Input[str] src6: IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
        :param pulumi.Input[str] status: Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
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
    def addr_type(self) -> Optional[pulumi.Input[str]]:
        """
        DNS translation type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        """
        return pulumi.get(self, "addr_type")

    @addr_type.setter
    def addr_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_type", value)

    @property
    @pulumi.getter
    def dst(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
        """
        return pulumi.get(self, "dst")

    @dst.setter
    def dst(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst", value)

    @property
    @pulumi.getter
    def dst6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
        """
        return pulumi.get(self, "dst6")

    @dst6.setter
    def dst6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst6", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[int]]:
        """
        If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def src(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
        """
        return pulumi.get(self, "src")

    @src.setter
    def src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src", value)

    @property
    @pulumi.getter
    def src6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
        """
        return pulumi.get(self, "src6")

    @src6.setter
    def src6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src6", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this DNS translation entry. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class ProfileDomainFilterArgs:
    def __init__(__self__, *,
                 domain_filter_table: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] domain_filter_table: DNS domain filter table ID.
        """
        if domain_filter_table is not None:
            pulumi.set(__self__, "domain_filter_table", domain_filter_table)

    @property
    @pulumi.getter(name="domainFilterTable")
    def domain_filter_table(self) -> Optional[pulumi.Input[int]]:
        """
        DNS domain filter table ID.
        """
        return pulumi.get(self, "domain_filter_table")

    @domain_filter_table.setter
    def domain_filter_table(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "domain_filter_table", value)


@pulumi.input_type
class ProfileExternalIpBlocklistArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: External domain block list name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        External domain block list name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ProfileFtgdDnsArgs:
    def __init__(__self__, *,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileFtgdDnsFilterArgs']]]] = None,
                 options: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ProfileFtgdDnsFilterArgs']]] filters: FortiGuard DNS domain filters. The structure of `filters` block is documented below.
        :param pulumi.Input[str] options: FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
        """
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if options is not None:
            pulumi.set(__self__, "options", options)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileFtgdDnsFilterArgs']]]]:
        """
        FortiGuard DNS domain filters. The structure of `filters` block is documented below.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileFtgdDnsFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "options", value)


@pulumi.input_type
class ProfileFtgdDnsFilterArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[int]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
        :param pulumi.Input[int] category: Category number.
        :param pulumi.Input[int] id: ID number.
        :param pulumi.Input[str] log: Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
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
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take for DNS requests matching the category. Valid values: `block`, `monitor`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[int]]:
        """
        Category number.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID number.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DNS filter logging for this DNS profile. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

