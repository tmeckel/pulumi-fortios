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
from ._inputs import *

__all__ = ['FortigateprofileArgs', 'Fortigateprofile']

@pulumi.input_type
class FortigateprofileArgs:
    def __init__(__self__, *,
                 extension: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 lan_extension: Optional[pulumi.Input['FortigateprofileLanExtensionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fortigateprofile resource.
        :param pulumi.Input[str] extension: Extension option. Valid values: `lan-extension`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input['FortigateprofileLanExtensionArgs'] lan_extension: FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        :param pulumi.Input[str] name: FortiGate connector profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if extension is not None:
            pulumi.set(__self__, "extension", extension)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if lan_extension is not None:
            pulumi.set(__self__, "lan_extension", lan_extension)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def extension(self) -> Optional[pulumi.Input[str]]:
        """
        Extension option. Valid values: `lan-extension`.
        """
        return pulumi.get(self, "extension")

    @extension.setter
    def extension(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extension", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="lanExtension")
    def lan_extension(self) -> Optional[pulumi.Input['FortigateprofileLanExtensionArgs']]:
        """
        FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        """
        return pulumi.get(self, "lan_extension")

    @lan_extension.setter
    def lan_extension(self, value: Optional[pulumi.Input['FortigateprofileLanExtensionArgs']]):
        pulumi.set(self, "lan_extension", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate connector profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FortigateprofileState:
    def __init__(__self__, *,
                 extension: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 lan_extension: Optional[pulumi.Input['FortigateprofileLanExtensionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Fortigateprofile resources.
        :param pulumi.Input[str] extension: Extension option. Valid values: `lan-extension`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input['FortigateprofileLanExtensionArgs'] lan_extension: FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        :param pulumi.Input[str] name: FortiGate connector profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if extension is not None:
            pulumi.set(__self__, "extension", extension)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if lan_extension is not None:
            pulumi.set(__self__, "lan_extension", lan_extension)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def extension(self) -> Optional[pulumi.Input[str]]:
        """
        Extension option. Valid values: `lan-extension`.
        """
        return pulumi.get(self, "extension")

    @extension.setter
    def extension(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extension", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="lanExtension")
    def lan_extension(self) -> Optional[pulumi.Input['FortigateprofileLanExtensionArgs']]:
        """
        FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        """
        return pulumi.get(self, "lan_extension")

    @lan_extension.setter
    def lan_extension(self, value: Optional[pulumi.Input['FortigateprofileLanExtensionArgs']]):
        pulumi.set(self, "lan_extension", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiGate connector profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Fortigateprofile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extension: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 lan_extension: Optional[pulumi.Input[pulumi.InputType['FortigateprofileLanExtensionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        FortiGate connector profile configuration. Applies to FortiOS Version `>= 7.2.1`.

        ## Import

        ExtensionController FortigateProfile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigateprofile:Fortigateprofile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigateprofile:Fortigateprofile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] extension: Extension option. Valid values: `lan-extension`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[pulumi.InputType['FortigateprofileLanExtensionArgs']] lan_extension: FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        :param pulumi.Input[str] name: FortiGate connector profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortigateprofileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        FortiGate connector profile configuration. Applies to FortiOS Version `>= 7.2.1`.

        ## Import

        ExtensionController FortigateProfile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigateprofile:Fortigateprofile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:extensioncontroller/fortigateprofile:Fortigateprofile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FortigateprofileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortigateprofileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extension: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 lan_extension: Optional[pulumi.Input[pulumi.InputType['FortigateprofileLanExtensionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortigateprofileArgs.__new__(FortigateprofileArgs)

            __props__.__dict__["extension"] = extension
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["lan_extension"] = lan_extension
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Fortigateprofile, __self__).__init__(
            'fortios:extensioncontroller/fortigateprofile:Fortigateprofile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            extension: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            lan_extension: Optional[pulumi.Input[pulumi.InputType['FortigateprofileLanExtensionArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Fortigateprofile':
        """
        Get an existing Fortigateprofile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] extension: Extension option. Valid values: `lan-extension`.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[pulumi.InputType['FortigateprofileLanExtensionArgs']] lan_extension: FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        :param pulumi.Input[str] name: FortiGate connector profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortigateprofileState.__new__(_FortigateprofileState)

        __props__.__dict__["extension"] = extension
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["lan_extension"] = lan_extension
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Fortigateprofile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def extension(self) -> pulumi.Output[str]:
        """
        Extension option. Valid values: `lan-extension`.
        """
        return pulumi.get(self, "extension")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="lanExtension")
    def lan_extension(self) -> pulumi.Output['outputs.FortigateprofileLanExtension']:
        """
        FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
        """
        return pulumi.get(self, "lan_extension")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FortiGate connector profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

