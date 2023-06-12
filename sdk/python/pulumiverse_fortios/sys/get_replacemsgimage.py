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
    'GetReplacemsgimageResult',
    'AwaitableGetReplacemsgimageResult',
    'get_replacemsgimage',
    'get_replacemsgimage_output',
]

@pulumi.output_type
class GetReplacemsgimageResult:
    """
    A collection of values returned by getReplacemsgimage.
    """
    def __init__(__self__, id=None, image_base64=None, image_type=None, name=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_base64 and not isinstance(image_base64, str):
            raise TypeError("Expected argument 'image_base64' to be a str")
        pulumi.set(__self__, "image_base64", image_base64)
        if image_type and not isinstance(image_type, str):
            raise TypeError("Expected argument 'image_type' to be a str")
        pulumi.set(__self__, "image_type", image_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageBase64")
    def image_base64(self) -> str:
        """
        Image data.
        """
        return pulumi.get(self, "image_base64")

    @property
    @pulumi.getter(name="imageType")
    def image_type(self) -> str:
        """
        Image type.
        """
        return pulumi.get(self, "image_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Image name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetReplacemsgimageResult(GetReplacemsgimageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReplacemsgimageResult(
            id=self.id,
            image_base64=self.image_base64,
            image_type=self.image_type,
            name=self.name,
            vdomparam=self.vdomparam)


def get_replacemsgimage(name: Optional[str] = None,
                        vdomparam: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReplacemsgimageResult:
    """
    Use this data source to get information on an fortios system replacemsgimage


    :param str name: Specify the name of the desired system replacemsgimage.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:sys/getReplacemsgimage:getReplacemsgimage', __args__, opts=opts, typ=GetReplacemsgimageResult).value

    return AwaitableGetReplacemsgimageResult(
        id=__ret__.id,
        image_base64=__ret__.image_base64,
        image_type=__ret__.image_type,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_replacemsgimage)
def get_replacemsgimage_output(name: Optional[pulumi.Input[str]] = None,
                               vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReplacemsgimageResult]:
    """
    Use this data source to get information on an fortios system replacemsgimage


    :param str name: Specify the name of the desired system replacemsgimage.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...