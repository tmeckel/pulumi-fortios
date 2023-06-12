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
    'PolicyAppCategoryArgs',
    'PolicyAppGroupArgs',
    'PolicyApplicationArgs',
    'PolicyDstaddr4Args',
    'PolicyDstaddr6Args',
    'PolicyDstintfArgs',
    'PolicyFssoGroupArgs',
    'PolicyGroupArgs',
    'PolicyInternetServiceCustomArgs',
    'PolicyInternetServiceCustomGroupArgs',
    'PolicyInternetServiceGroupArgs',
    'PolicyInternetServiceIdArgs',
    'PolicyInternetServiceNameArgs',
    'PolicyInternetServiceSrcCustomArgs',
    'PolicyInternetServiceSrcCustomGroupArgs',
    'PolicyInternetServiceSrcGroupArgs',
    'PolicyInternetServiceSrcIdArgs',
    'PolicyInternetServiceSrcNameArgs',
    'PolicyPoolname4Args',
    'PolicyPoolname6Args',
    'PolicyServiceArgs',
    'PolicySrcaddr4Args',
    'PolicySrcaddr6Args',
    'PolicySrcintfArgs',
    'PolicyUrlCategoryArgs',
    'PolicyUserArgs',
]

@pulumi.input_type
class PolicyAppCategoryArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] id: Category IDs.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Category IDs.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class PolicyAppGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Application group names.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Application group names.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyApplicationArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] id: Application IDs.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Application IDs.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class PolicyDstaddr4Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyDstaddr6Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyDstintfArgs:
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
class PolicyFssoGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Names of FSSO groups.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Names of FSSO groups.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceCustomArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Custom Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Custom Internet Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceCustomGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Custom Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Custom Internet Service group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Internet Service group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceIdArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] id: Internet Service ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class PolicyInternetServiceNameArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceSrcCustomArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Custom Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Custom Internet Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceSrcCustomGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Custom Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Custom Internet Service group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceSrcGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Internet Service group name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Internet Service group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyInternetServiceSrcIdArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] id: Internet Service ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class PolicyInternetServiceSrcNameArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Internet Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyPoolname4Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyPoolname6Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyServiceArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Service name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicySrcaddr4Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicySrcaddr6Args:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Policy name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicySrcintfArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Interface name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicyUrlCategoryArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] id: URL category ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        URL category ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class PolicyUserArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: IPv6 pool name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 pool name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


