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
    'FilterFreeStyleArgs',
    'OverridefilterFreeStyleArgs',
    'OverridesettingCustomFieldNameArgs',
    'SettingCustomFieldNameArgs',
]

@pulumi.input_type
class FilterFreeStyleArgs:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] category: Log category.
        :param pulumi.Input[str] filter: Free style filter string.
        :param pulumi.Input[str] filter_type: Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        :param pulumi.Input[int] id: Entry ID.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if filter_type is not None:
            pulumi.set(__self__, "filter_type", filter_type)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Log category.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Free style filter string.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        """
        Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class OverridefilterFreeStyleArgs:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 filter_type: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] category: Log category.
        :param pulumi.Input[str] filter: Free style filter string.
        :param pulumi.Input[str] filter_type: Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        :param pulumi.Input[int] id: Entry ID.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if filter_type is not None:
            pulumi.set(__self__, "filter_type", filter_type)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Log category.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Free style filter string.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="filterType")
    def filter_type(self) -> Optional[pulumi.Input[str]]:
        """
        Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
        """
        return pulumi.get(self, "filter_type")

    @filter_type.setter
    def filter_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class OverridesettingCustomFieldNameArgs:
    def __init__(__self__, *,
                 custom: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] custom: Field custom name.
        :param pulumi.Input[int] id: Entry ID.
        :param pulumi.Input[str] name: Field name.
        """
        if custom is not None:
            pulumi.set(__self__, "custom", custom)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def custom(self) -> Optional[pulumi.Input[str]]:
        """
        Field custom name.
        """
        return pulumi.get(self, "custom")

    @custom.setter
    def custom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SettingCustomFieldNameArgs:
    def __init__(__self__, *,
                 custom: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] custom: Field custom name.
        :param pulumi.Input[int] id: Entry ID.
        :param pulumi.Input[str] name: Field name.
        """
        if custom is not None:
            pulumi.set(__self__, "custom", custom)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def custom(self) -> Optional[pulumi.Input[str]]:
        """
        Field custom name.
        """
        return pulumi.get(self, "custom")

    @custom.setter
    def custom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Entry ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


