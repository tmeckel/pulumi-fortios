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

__all__ = ['SsoadminArgs', 'Ssoadmin']

@pulumi.input_type
class SsoadminArgs:
    def __init__(__self__, *,
                 accprofile: pulumi.Input[str],
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gui_ignore_release_overview_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]]] = None):
        """
        The set of arguments for constructing a Ssoadmin resource.
        :param pulumi.Input[str] accprofile: SSO admin user access profile.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] gui_ignore_release_overview_version: The FortiOS version to ignore release overview prompt for.
        :param pulumi.Input[str] name: SSO admin name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]] vdoms: Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        pulumi.set(__self__, "accprofile", accprofile)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if gui_ignore_release_overview_version is not None:
            pulumi.set(__self__, "gui_ignore_release_overview_version", gui_ignore_release_overview_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Input[str]:
        """
        SSO admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @accprofile.setter
    def accprofile(self, value: pulumi.Input[str]):
        pulumi.set(self, "accprofile", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="guiIgnoreReleaseOverviewVersion")
    def gui_ignore_release_overview_version(self) -> Optional[pulumi.Input[str]]:
        """
        The FortiOS version to ignore release overview prompt for.
        """
        return pulumi.get(self, "gui_ignore_release_overview_version")

    @gui_ignore_release_overview_version.setter
    def gui_ignore_release_overview_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gui_ignore_release_overview_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SSO admin name.
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

    @property
    @pulumi.getter
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]]]:
        """
        Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]]]):
        pulumi.set(self, "vdoms", value)


@pulumi.input_type
class _SsoadminState:
    def __init__(__self__, *,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gui_ignore_release_overview_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]]] = None):
        """
        Input properties used for looking up and filtering Ssoadmin resources.
        :param pulumi.Input[str] accprofile: SSO admin user access profile.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] gui_ignore_release_overview_version: The FortiOS version to ignore release overview prompt for.
        :param pulumi.Input[str] name: SSO admin name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]] vdoms: Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        if accprofile is not None:
            pulumi.set(__self__, "accprofile", accprofile)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if gui_ignore_release_overview_version is not None:
            pulumi.set(__self__, "gui_ignore_release_overview_version", gui_ignore_release_overview_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def accprofile(self) -> Optional[pulumi.Input[str]]:
        """
        SSO admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @accprofile.setter
    def accprofile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accprofile", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="guiIgnoreReleaseOverviewVersion")
    def gui_ignore_release_overview_version(self) -> Optional[pulumi.Input[str]]:
        """
        The FortiOS version to ignore release overview prompt for.
        """
        return pulumi.get(self, "gui_ignore_release_overview_version")

    @gui_ignore_release_overview_version.setter
    def gui_ignore_release_overview_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gui_ignore_release_overview_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SSO admin name.
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

    @property
    @pulumi.getter
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]]]:
        """
        Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SsoadminVdomArgs']]]]):
        pulumi.set(self, "vdoms", value)


class Ssoadmin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gui_ignore_release_overview_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SsoadminVdomArgs']]]]] = None,
                 __props__=None):
        """
        Configure SSO admin users. Applies to FortiOS Version `>= 6.2.4`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Ssoadmin("trname",
            accprofile="super_admin",
            vdoms=[fortios.sys.SsoadminVdomArgs(
                name="root",
            )])
        ```

        ## Import

        System SsoAdmin can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/ssoadmin:Ssoadmin labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/ssoadmin:Ssoadmin labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: SSO admin user access profile.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] gui_ignore_release_overview_version: The FortiOS version to ignore release overview prompt for.
        :param pulumi.Input[str] name: SSO admin name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SsoadminVdomArgs']]]] vdoms: Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SsoadminArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure SSO admin users. Applies to FortiOS Version `>= 6.2.4`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Ssoadmin("trname",
            accprofile="super_admin",
            vdoms=[fortios.sys.SsoadminVdomArgs(
                name="root",
            )])
        ```

        ## Import

        System SsoAdmin can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/ssoadmin:Ssoadmin labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/ssoadmin:Ssoadmin labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SsoadminArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SsoadminArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 gui_ignore_release_overview_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SsoadminVdomArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SsoadminArgs.__new__(SsoadminArgs)

            if accprofile is None and not opts.urn:
                raise TypeError("Missing required property 'accprofile'")
            __props__.__dict__["accprofile"] = accprofile
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["gui_ignore_release_overview_version"] = gui_ignore_release_overview_version
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vdoms"] = vdoms
        super(Ssoadmin, __self__).__init__(
            'fortios:sys/ssoadmin:Ssoadmin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accprofile: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            gui_ignore_release_overview_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SsoadminVdomArgs']]]]] = None) -> 'Ssoadmin':
        """
        Get an existing Ssoadmin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: SSO admin user access profile.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] gui_ignore_release_overview_version: The FortiOS version to ignore release overview prompt for.
        :param pulumi.Input[str] name: SSO admin name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SsoadminVdomArgs']]]] vdoms: Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SsoadminState.__new__(_SsoadminState)

        __props__.__dict__["accprofile"] = accprofile
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["gui_ignore_release_overview_version"] = gui_ignore_release_overview_version
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vdoms"] = vdoms
        return Ssoadmin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Output[str]:
        """
        SSO admin user access profile.
        """
        return pulumi.get(self, "accprofile")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="guiIgnoreReleaseOverviewVersion")
    def gui_ignore_release_overview_version(self) -> pulumi.Output[str]:
        """
        The FortiOS version to ignore release overview prompt for.
        """
        return pulumi.get(self, "gui_ignore_release_overview_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        SSO admin name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vdoms(self) -> pulumi.Output[Optional[Sequence['outputs.SsoadminVdom']]]:
        """
        Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")
