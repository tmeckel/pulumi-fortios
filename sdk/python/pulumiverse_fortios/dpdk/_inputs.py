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
    'GlobalInterfaceArgs',
]

@pulumi.input_type
class GlobalInterfaceArgs:
    def __init__(__self__, *,
                 interface_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] interface_name: Physical interface name.
        """
        if interface_name is not None:
            pulumi.set(__self__, "interface_name", interface_name)

    @property
    @pulumi.getter(name="interfaceName")
    def interface_name(self) -> Optional[pulumi.Input[str]]:
        """
        Physical interface name.
        """
        return pulumi.get(self, "interface_name")

    @interface_name.setter
    def interface_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_name", value)


