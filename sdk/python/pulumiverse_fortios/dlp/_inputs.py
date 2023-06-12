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
    'DictionaryEntryArgs',
    'FilepatternEntryArgs',
    'ProfileRuleArgs',
    'ProfileRuleSensitivityArgs',
    'ProfileRuleSensorArgs',
    'SensorEntryArgs',
    'SensorFilterArgs',
    'SensorFilterFpSensitivityArgs',
    'SensorFilterSensitivityArgs',
]

@pulumi.input_type
class DictionaryEntryArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 ignore_case: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repeat: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] ignore_case: Enable/disable ignore case. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pattern: Pattern to match.
        :param pulumi.Input[str] repeat: Enable/disable repeat match. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable this pattern. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: Pattern type to match.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ignore_case is not None:
            pulumi.set(__self__, "ignore_case", ignore_case)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if repeat is not None:
            pulumi.set(__self__, "repeat", repeat)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Optional comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

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
    @pulumi.getter(name="ignoreCase")
    def ignore_case(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable ignore case. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ignore_case")

    @ignore_case.setter
    def ignore_case(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ignore_case", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Pattern to match.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def repeat(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable repeat match. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "repeat")

    @repeat.setter
    def repeat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this pattern. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Pattern type to match.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class FilepatternEntryArgs:
    def __init__(__self__, *,
                 file_type: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] file_type: Select a file type.
        :param pulumi.Input[str] filter_type: Filter by file name pattern or by file type. Valid values: `pattern`, `type`.
        :param pulumi.Input[str] pattern: Add a file name pattern.
        """
        if file_type is not None:
            pulumi.set(__self__, "file_type", file_type)
        if filter_type is not None:
            pulumi.set(__self__, "filter_type", filter_type)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> Optional[pulumi.Input[str]]:
        """
        Select a file type.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        """
        Filter by file name pattern or by file type. Valid values: `pattern`, `type`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Add a file name pattern.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)


@pulumi.input_type
class ProfileRuleArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 archive: Optional[pulumi.Input[str]] = None,
                 expiry: Optional[pulumi.Input[str]] = None,
                 file_size: Optional[pulumi.Input[int]] = None,
                 file_type: Optional[pulumi.Input[int]] = None,
                 filter_by: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 match_percentage: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 proto: Optional[pulumi.Input[str]] = None,
                 sensitivities: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensitivityArgs']]]] = None,
                 sensors: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensorArgs']]]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: Action to take with content that this DLP profile matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
        :param pulumi.Input[str] archive: Enable/disable DLP archiving. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] expiry: Quarantine duration in days, hours, minutes (format = dddhhmm).
        :param pulumi.Input[int] file_size: Match files this size or larger (0 - 4294967295 kbytes).
        :param pulumi.Input[int] file_type: Select the number of a DLP file pattern table to match.
        :param pulumi.Input[str] filter_by: Select the type of content to match. Valid values: `sensor`, `mip`, `fingerprint`, `encrypted`, `none`.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] label: MIP label dictionary.
        :param pulumi.Input[int] match_percentage: Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
        :param pulumi.Input[str] name: Filter name.
        :param pulumi.Input[str] proto: Check messages or files over one or more of these protocols. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensitivityArgs']]] sensitivities: Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensorArgs']]] sensors: Select DLP sensors. The structure of `sensor` block is documented below.
        :param pulumi.Input[str] severity: Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] type: Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if archive is not None:
            pulumi.set(__self__, "archive", archive)
        if expiry is not None:
            pulumi.set(__self__, "expiry", expiry)
        if file_size is not None:
            pulumi.set(__self__, "file_size", file_size)
        if file_type is not None:
            pulumi.set(__self__, "file_type", file_type)
        if filter_by is not None:
            pulumi.set(__self__, "filter_by", filter_by)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if match_percentage is not None:
            pulumi.set(__self__, "match_percentage", match_percentage)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if proto is not None:
            pulumi.set(__self__, "proto", proto)
        if sensitivities is not None:
            pulumi.set(__self__, "sensitivities", sensitivities)
        if sensors is not None:
            pulumi.set(__self__, "sensors", sensors)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take with content that this DLP profile matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def archive(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DLP archiving. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "archive")

    @archive.setter
    def archive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "archive", value)

    @property
    @pulumi.getter
    def expiry(self) -> Optional[pulumi.Input[str]]:
        """
        Quarantine duration in days, hours, minutes (format = dddhhmm).
        """
        return pulumi.get(self, "expiry")

    @expiry.setter
    def expiry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry", value)

    @property
    @pulumi.getter(name="fileSize")
    def file_size(self) -> Optional[pulumi.Input[int]]:
        """
        Match files this size or larger (0 - 4294967295 kbytes).
        """
        return pulumi.get(self, "file_size")

    @file_size.setter
    def file_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "file_size", value)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> Optional[pulumi.Input[int]]:
        """
        Select the number of a DLP file pattern table to match.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter(name="filterBy")
    def filter_by(self) -> Optional[pulumi.Input[str]]:
        """
        Select the type of content to match. Valid values: `sensor`, `mip`, `fingerprint`, `encrypted`, `none`.
        """
        return pulumi.get(self, "filter_by")

    @filter_by.setter
    def filter_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_by", value)

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
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        MIP label dictionary.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="matchPercentage")
    def match_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
        """
        return pulumi.get(self, "match_percentage")

    @match_percentage.setter
    def match_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_percentage", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Filter name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def proto(self) -> Optional[pulumi.Input[str]]:
        """
        Check messages or files over one or more of these protocols. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "proto")

    @proto.setter
    def proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proto", value)

    @property
    @pulumi.getter
    def sensitivities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensitivityArgs']]]]:
        """
        Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
        """
        return pulumi.get(self, "sensitivities")

    @sensitivities.setter
    def sensitivities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensitivityArgs']]]]):
        pulumi.set(self, "sensitivities", value)

    @property
    @pulumi.getter
    def sensors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensorArgs']]]]:
        """
        Select DLP sensors. The structure of `sensor` block is documented below.
        """
        return pulumi.get(self, "sensors")

    @sensors.setter
    def sensors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleSensorArgs']]]]):
        pulumi.set(self, "sensors", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProfileRuleSensitivityArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Select a DLP sensitivity.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Select a DLP sensitivity.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ProfileRuleSensorArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Address name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Address name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SensorEntryArgs:
    def __init__(__self__, *,
                 count: Optional[pulumi.Input[int]] = None,
                 dictionary: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] count: Count of dictionary matches to trigger sensor entry match (Dictionary might not be able to trigger more than once based on its 'repeat' option, 1 - 255, default = 1).
        :param pulumi.Input[str] dictionary: Select a DLP dictionary.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] status: Enable/disable this entry. Valid values: `enable`, `disable`.
        """
        if count is not None:
            pulumi.set(__self__, "count", count)
        if dictionary is not None:
            pulumi.set(__self__, "dictionary", dictionary)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def count(self) -> Optional[pulumi.Input[int]]:
        """
        Count of dictionary matches to trigger sensor entry match (Dictionary might not be able to trigger more than once based on its 'repeat' option, 1 - 255, default = 1).
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def dictionary(self) -> Optional[pulumi.Input[str]]:
        """
        Select a DLP dictionary.
        """
        return pulumi.get(self, "dictionary")

    @dictionary.setter
    def dictionary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dictionary", value)

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
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this entry. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class SensorFilterArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 archive: Optional[pulumi.Input[str]] = None,
                 company_identifier: Optional[pulumi.Input[str]] = None,
                 expiry: Optional[pulumi.Input[str]] = None,
                 file_size: Optional[pulumi.Input[int]] = None,
                 file_type: Optional[pulumi.Input[int]] = None,
                 filter_by: Optional[pulumi.Input[str]] = None,
                 fp_sensitivities: Optional[pulumi.Input[Sequence[pulumi.Input['SensorFilterFpSensitivityArgs']]]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 match_percentage: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 proto: Optional[pulumi.Input[str]] = None,
                 regexp: Optional[pulumi.Input[str]] = None,
                 sensitivities: Optional[pulumi.Input[Sequence[pulumi.Input['SensorFilterSensitivityArgs']]]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action: Action to take with content that this DLP sensor matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
        :param pulumi.Input[str] archive: Enable/disable DLP archiving. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] company_identifier: Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
        :param pulumi.Input[str] expiry: Quarantine duration in days, hours, minutes format (dddhhmm).
        :param pulumi.Input[int] file_size: Match files this size or larger (0 - 4294967295 kbytes).
        :param pulumi.Input[int] file_type: Select the number of a DLP file pattern table to match.
        :param pulumi.Input[str] filter_by: Select the type of content to match.
        :param pulumi.Input[Sequence[pulumi.Input['SensorFilterFpSensitivityArgs']]] fp_sensitivities: Select a DLP file pattern sensitivity to match. The structure of `fp_sensitivity` block is documented below.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[int] match_percentage: Percentage of fingerprints in the fingerprint databases designated with the selected fp-sensitivity to match.
        :param pulumi.Input[str] name: Filter name.
        :param pulumi.Input[str] proto: Check messages or files over one or more of these protocols.
        :param pulumi.Input[str] regexp: Enter a regular expression to match (max. 255 characters).
        :param pulumi.Input[Sequence[pulumi.Input['SensorFilterSensitivityArgs']]] sensitivities: Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
        :param pulumi.Input[str] severity: Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
        :param pulumi.Input[str] type: Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).  Valid values: `file`, `message`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if archive is not None:
            pulumi.set(__self__, "archive", archive)
        if company_identifier is not None:
            pulumi.set(__self__, "company_identifier", company_identifier)
        if expiry is not None:
            pulumi.set(__self__, "expiry", expiry)
        if file_size is not None:
            pulumi.set(__self__, "file_size", file_size)
        if file_type is not None:
            pulumi.set(__self__, "file_type", file_type)
        if filter_by is not None:
            pulumi.set(__self__, "filter_by", filter_by)
        if fp_sensitivities is not None:
            pulumi.set(__self__, "fp_sensitivities", fp_sensitivities)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if match_percentage is not None:
            pulumi.set(__self__, "match_percentage", match_percentage)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if proto is not None:
            pulumi.set(__self__, "proto", proto)
        if regexp is not None:
            pulumi.set(__self__, "regexp", regexp)
        if sensitivities is not None:
            pulumi.set(__self__, "sensitivities", sensitivities)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take with content that this DLP sensor matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def archive(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DLP archiving. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "archive")

    @archive.setter
    def archive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "archive", value)

    @property
    @pulumi.getter(name="companyIdentifier")
    def company_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
        """
        return pulumi.get(self, "company_identifier")

    @company_identifier.setter
    def company_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "company_identifier", value)

    @property
    @pulumi.getter
    def expiry(self) -> Optional[pulumi.Input[str]]:
        """
        Quarantine duration in days, hours, minutes format (dddhhmm).
        """
        return pulumi.get(self, "expiry")

    @expiry.setter
    def expiry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry", value)

    @property
    @pulumi.getter(name="fileSize")
    def file_size(self) -> Optional[pulumi.Input[int]]:
        """
        Match files this size or larger (0 - 4294967295 kbytes).
        """
        return pulumi.get(self, "file_size")

    @file_size.setter
    def file_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "file_size", value)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> Optional[pulumi.Input[int]]:
        """
        Select the number of a DLP file pattern table to match.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter(name="filterBy")
    def filter_by(self) -> Optional[pulumi.Input[str]]:
        """
        Select the type of content to match.
        """
        return pulumi.get(self, "filter_by")

    @filter_by.setter
    def filter_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_by", value)

    @property
    @pulumi.getter(name="fpSensitivities")
    def fp_sensitivities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SensorFilterFpSensitivityArgs']]]]:
        """
        Select a DLP file pattern sensitivity to match. The structure of `fp_sensitivity` block is documented below.
        """
        return pulumi.get(self, "fp_sensitivities")

    @fp_sensitivities.setter
    def fp_sensitivities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SensorFilterFpSensitivityArgs']]]]):
        pulumi.set(self, "fp_sensitivities", value)

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
    @pulumi.getter(name="matchPercentage")
    def match_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Percentage of fingerprints in the fingerprint databases designated with the selected fp-sensitivity to match.
        """
        return pulumi.get(self, "match_percentage")

    @match_percentage.setter
    def match_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_percentage", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Filter name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def proto(self) -> Optional[pulumi.Input[str]]:
        """
        Check messages or files over one or more of these protocols.
        """
        return pulumi.get(self, "proto")

    @proto.setter
    def proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proto", value)

    @property
    @pulumi.getter
    def regexp(self) -> Optional[pulumi.Input[str]]:
        """
        Enter a regular expression to match (max. 255 characters).
        """
        return pulumi.get(self, "regexp")

    @regexp.setter
    def regexp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "regexp", value)

    @property
    @pulumi.getter
    def sensitivities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SensorFilterSensitivityArgs']]]]:
        """
        Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
        """
        return pulumi.get(self, "sensitivities")

    @sensitivities.setter
    def sensitivities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SensorFilterSensitivityArgs']]]]):
        pulumi.set(self, "sensitivities", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).  Valid values: `file`, `message`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SensorFilterFpSensitivityArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Select a DLP sensitivity.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Select a DLP sensitivity.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SensorFilterSensitivityArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Select a DLP sensitivity.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Select a DLP sensitivity.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


