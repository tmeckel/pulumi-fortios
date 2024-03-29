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
    'BwlEntry',
    'BwordEntry',
    'DnsblEntry',
    'IptrustEntry',
    'MheaderEntry',
    'ProfileGmail',
    'ProfileImap',
    'ProfileMapi',
    'ProfileMsnHotmail',
    'ProfilePop3',
    'ProfileSmtp',
    'ProfileYahooMail',
]

@pulumi.output_type
class BwlEntry(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addrType":
            suggest = "addr_type"
        elif key == "emailPattern":
            suggest = "email_pattern"
        elif key == "ip4Subnet":
            suggest = "ip4_subnet"
        elif key == "ip6Subnet":
            suggest = "ip6_subnet"
        elif key == "patternType":
            suggest = "pattern_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BwlEntry. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BwlEntry.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BwlEntry.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 addr_type: Optional[str] = None,
                 email_pattern: Optional[str] = None,
                 id: Optional[int] = None,
                 ip4_subnet: Optional[str] = None,
                 ip6_subnet: Optional[str] = None,
                 pattern_type: Optional[str] = None,
                 status: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param str action: Reject, mark as spam or good email. Valid values: `reject`, `spam`, `clear`.
        :param str addr_type: IP address type. Valid values: `ipv4`, `ipv6`.
        :param str email_pattern: Email address pattern.
        :param int id: Entry ID.
        :param str ip4_subnet: IPv4 network address/subnet mask bits.
        :param str ip6_subnet: IPv6 network address/subnet mask bits.
        :param str pattern_type: Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        :param str status: Enable/disable status. Valid values: `enable`, `disable`.
        :param str type: Entry type. Valid values: `ip`, `email`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if addr_type is not None:
            pulumi.set(__self__, "addr_type", addr_type)
        if email_pattern is not None:
            pulumi.set(__self__, "email_pattern", email_pattern)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip4_subnet is not None:
            pulumi.set(__self__, "ip4_subnet", ip4_subnet)
        if ip6_subnet is not None:
            pulumi.set(__self__, "ip6_subnet", ip6_subnet)
        if pattern_type is not None:
            pulumi.set(__self__, "pattern_type", pattern_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Reject, mark as spam or good email. Valid values: `reject`, `spam`, `clear`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> Optional[str]:
        """
        IP address type. Valid values: `ipv4`, `ipv6`.
        """
        return pulumi.get(self, "addr_type")

    @property
    @pulumi.getter(name="emailPattern")
    def email_pattern(self) -> Optional[str]:
        """
        Email address pattern.
        """
        return pulumi.get(self, "email_pattern")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ip4Subnet")
    def ip4_subnet(self) -> Optional[str]:
        """
        IPv4 network address/subnet mask bits.
        """
        return pulumi.get(self, "ip4_subnet")

    @property
    @pulumi.getter(name="ip6Subnet")
    def ip6_subnet(self) -> Optional[str]:
        """
        IPv6 network address/subnet mask bits.
        """
        return pulumi.get(self, "ip6_subnet")

    @property
    @pulumi.getter(name="patternType")
    def pattern_type(self) -> Optional[str]:
        """
        Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        """
        return pulumi.get(self, "pattern_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable status. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Entry type. Valid values: `ip`, `email`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class BwordEntry(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "patternType":
            suggest = "pattern_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BwordEntry. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BwordEntry.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BwordEntry.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 id: Optional[int] = None,
                 language: Optional[str] = None,
                 pattern: Optional[str] = None,
                 pattern_type: Optional[str] = None,
                 score: Optional[int] = None,
                 status: Optional[str] = None,
                 where: Optional[str] = None):
        """
        :param str action: Mark spam or good. Valid values: `spam`, `clear`.
        :param int id: Banned word entry ID.
        :param str language: Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.
        :param str pattern: Pattern for the banned word.
        :param str pattern_type: Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        :param int score: Score value.
        :param str status: Enable/disable status. Valid values: `enable`, `disable`.
        :param str where: Component of the email to be scanned. Valid values: `subject`, `body`, `all`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if pattern_type is not None:
            pulumi.set(__self__, "pattern_type", pattern_type)
        if score is not None:
            pulumi.set(__self__, "score", score)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if where is not None:
            pulumi.set(__self__, "where", where)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Mark spam or good. Valid values: `spam`, `clear`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Banned word entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def language(self) -> Optional[str]:
        """
        Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def pattern(self) -> Optional[str]:
        """
        Pattern for the banned word.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter(name="patternType")
    def pattern_type(self) -> Optional[str]:
        """
        Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        """
        return pulumi.get(self, "pattern_type")

    @property
    @pulumi.getter
    def score(self) -> Optional[int]:
        """
        Score value.
        """
        return pulumi.get(self, "score")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable status. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def where(self) -> Optional[str]:
        """
        Component of the email to be scanned. Valid values: `subject`, `body`, `all`.
        """
        return pulumi.get(self, "where")


@pulumi.output_type
class DnsblEntry(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 id: Optional[int] = None,
                 server: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str action: Reject connection or mark as spam email. Valid values: `reject`, `spam`.
        :param int id: DNSBL/ORBL entry ID.
        :param str server: DNSBL or ORBL server name.
        :param str status: Enable/disable status. Valid values: `enable`, `disable`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Reject connection or mark as spam email. Valid values: `reject`, `spam`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        DNSBL/ORBL entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def server(self) -> Optional[str]:
        """
        DNSBL or ORBL server name.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable status. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class IptrustEntry(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addrType":
            suggest = "addr_type"
        elif key == "ip4Subnet":
            suggest = "ip4_subnet"
        elif key == "ip6Subnet":
            suggest = "ip6_subnet"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IptrustEntry. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IptrustEntry.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IptrustEntry.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 addr_type: Optional[str] = None,
                 id: Optional[int] = None,
                 ip4_subnet: Optional[str] = None,
                 ip6_subnet: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str addr_type: Type of address. Valid values: `ipv4`, `ipv6`.
        :param int id: Trusted IP entry ID.
        :param str ip4_subnet: IPv4 network address or network address/subnet mask bits.
        :param str ip6_subnet: IPv6 network address/subnet mask bits.
        :param str status: Enable/disable status. Valid values: `enable`, `disable`.
        """
        if addr_type is not None:
            pulumi.set(__self__, "addr_type", addr_type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip4_subnet is not None:
            pulumi.set(__self__, "ip4_subnet", ip4_subnet)
        if ip6_subnet is not None:
            pulumi.set(__self__, "ip6_subnet", ip6_subnet)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> Optional[str]:
        """
        Type of address. Valid values: `ipv4`, `ipv6`.
        """
        return pulumi.get(self, "addr_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Trusted IP entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ip4Subnet")
    def ip4_subnet(self) -> Optional[str]:
        """
        IPv4 network address or network address/subnet mask bits.
        """
        return pulumi.get(self, "ip4_subnet")

    @property
    @pulumi.getter(name="ip6Subnet")
    def ip6_subnet(self) -> Optional[str]:
        """
        IPv6 network address/subnet mask bits.
        """
        return pulumi.get(self, "ip6_subnet")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable status. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class MheaderEntry(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "patternType":
            suggest = "pattern_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MheaderEntry. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MheaderEntry.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MheaderEntry.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 fieldbody: Optional[str] = None,
                 fieldname: Optional[str] = None,
                 id: Optional[int] = None,
                 pattern_type: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str action: Mark spam or good. Valid values: `spam`, `clear`.
        :param str fieldbody: Pattern for the header field body.
        :param str fieldname: Pattern for header field name.
        :param int id: Mime header entry ID.
        :param str pattern_type: Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        :param str status: Enable/disable status. Valid values: `enable`, `disable`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if fieldbody is not None:
            pulumi.set(__self__, "fieldbody", fieldbody)
        if fieldname is not None:
            pulumi.set(__self__, "fieldname", fieldname)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if pattern_type is not None:
            pulumi.set(__self__, "pattern_type", pattern_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Mark spam or good. Valid values: `spam`, `clear`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def fieldbody(self) -> Optional[str]:
        """
        Pattern for the header field body.
        """
        return pulumi.get(self, "fieldbody")

    @property
    @pulumi.getter
    def fieldname(self) -> Optional[str]:
        """
        Pattern for header field name.
        """
        return pulumi.get(self, "fieldname")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        Mime header entry ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="patternType")
    def pattern_type(self) -> Optional[str]:
        """
        Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        """
        return pulumi.get(self, "pattern_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable/disable status. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ProfileGmail(dict):
    def __init__(__self__, *,
                 log: Optional[str] = None):
        """
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        """
        if log is not None:
            pulumi.set(__self__, "log", log)

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")


@pulumi.output_type
class ProfileImap(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "tagMsg":
            suggest = "tag_msg"
        elif key == "tagType":
            suggest = "tag_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileImap. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileImap.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileImap.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 log: Optional[str] = None,
                 tag_msg: Optional[str] = None,
                 tag_type: Optional[str] = None):
        """
        :param str action: Action for spam email. Valid values: `pass`, `tag`.
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        :param str tag_msg: Subject text or header added to spam email.
        :param str tag_type: Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if tag_msg is not None:
            pulumi.set(__self__, "tag_msg", tag_msg)
        if tag_type is not None:
            pulumi.set(__self__, "tag_type", tag_type)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action for spam email. Valid values: `pass`, `tag`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter(name="tagMsg")
    def tag_msg(self) -> Optional[str]:
        """
        Subject text or header added to spam email.
        """
        return pulumi.get(self, "tag_msg")

    @property
    @pulumi.getter(name="tagType")
    def tag_type(self) -> Optional[str]:
        """
        Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        """
        return pulumi.get(self, "tag_type")


@pulumi.output_type
class ProfileMapi(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 log: Optional[str] = None):
        """
        :param str action: Action for spam email. Valid values: `pass`, `discard`.
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if log is not None:
            pulumi.set(__self__, "log", log)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action for spam email. Valid values: `pass`, `discard`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")


@pulumi.output_type
class ProfileMsnHotmail(dict):
    def __init__(__self__, *,
                 log: Optional[str] = None):
        """
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        """
        if log is not None:
            pulumi.set(__self__, "log", log)

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")


@pulumi.output_type
class ProfilePop3(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "tagMsg":
            suggest = "tag_msg"
        elif key == "tagType":
            suggest = "tag_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfilePop3. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfilePop3.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfilePop3.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 log: Optional[str] = None,
                 tag_msg: Optional[str] = None,
                 tag_type: Optional[str] = None):
        """
        :param str action: Action for spam email. Valid values: `pass`, `tag`.
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        :param str tag_msg: Subject text or header added to spam email.
               
               The `pop3` block supports:
        :param str tag_type: Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if tag_msg is not None:
            pulumi.set(__self__, "tag_msg", tag_msg)
        if tag_type is not None:
            pulumi.set(__self__, "tag_type", tag_type)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action for spam email. Valid values: `pass`, `tag`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter(name="tagMsg")
    def tag_msg(self) -> Optional[str]:
        """
        Subject text or header added to spam email.

        The `pop3` block supports:
        """
        return pulumi.get(self, "tag_msg")

    @property
    @pulumi.getter(name="tagType")
    def tag_type(self) -> Optional[str]:
        """
        Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        """
        return pulumi.get(self, "tag_type")


@pulumi.output_type
class ProfileSmtp(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "localOverride":
            suggest = "local_override"
        elif key == "tagMsg":
            suggest = "tag_msg"
        elif key == "tagType":
            suggest = "tag_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileSmtp. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileSmtp.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileSmtp.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 hdrip: Optional[str] = None,
                 local_override: Optional[str] = None,
                 log: Optional[str] = None,
                 tag_msg: Optional[str] = None,
                 tag_type: Optional[str] = None):
        """
        :param str action: Action for spam email. Valid values: `pass`, `tag`, `discard`.
        :param str hdrip: Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.
        :param str local_override: Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        :param str tag_msg: Subject text or header added to spam email.
        :param str tag_type: Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if hdrip is not None:
            pulumi.set(__self__, "hdrip", hdrip)
        if local_override is not None:
            pulumi.set(__self__, "local_override", local_override)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if tag_msg is not None:
            pulumi.set(__self__, "tag_msg", tag_msg)
        if tag_type is not None:
            pulumi.set(__self__, "tag_type", tag_type)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        Action for spam email. Valid values: `pass`, `tag`, `discard`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def hdrip(self) -> Optional[str]:
        """
        Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "hdrip")

    @property
    @pulumi.getter(name="localOverride")
    def local_override(self) -> Optional[str]:
        """
        Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "local_override")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter(name="tagMsg")
    def tag_msg(self) -> Optional[str]:
        """
        Subject text or header added to spam email.
        """
        return pulumi.get(self, "tag_msg")

    @property
    @pulumi.getter(name="tagType")
    def tag_type(self) -> Optional[str]:
        """
        Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        """
        return pulumi.get(self, "tag_type")


@pulumi.output_type
class ProfileYahooMail(dict):
    def __init__(__self__, *,
                 log: Optional[str] = None):
        """
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        """
        if log is not None:
            pulumi.set(__self__, "log", log)

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")


