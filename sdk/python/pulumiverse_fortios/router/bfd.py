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

__all__ = ['BfdArgs', 'Bfd']

@pulumi.input_type
class BfdArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 multihop_templates: Optional[pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]]] = None,
                 neighbors: Optional[pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Bfd resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]] multihop_templates: BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]] neighbors: neighbor The structure of `neighbor` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if multihop_templates is not None:
            pulumi.set(__self__, "multihop_templates", multihop_templates)
        if neighbors is not None:
            pulumi.set(__self__, "neighbors", neighbors)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="multihopTemplates")
    def multihop_templates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]]]:
        """
        BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        """
        return pulumi.get(self, "multihop_templates")

    @multihop_templates.setter
    def multihop_templates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]]]):
        pulumi.set(self, "multihop_templates", value)

    @property
    @pulumi.getter
    def neighbors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]]]:
        """
        neighbor The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @neighbors.setter
    def neighbors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]]]):
        pulumi.set(self, "neighbors", value)

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
class _BfdState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 multihop_templates: Optional[pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]]] = None,
                 neighbors: Optional[pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Bfd resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]] multihop_templates: BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]] neighbors: neighbor The structure of `neighbor` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if multihop_templates is not None:
            pulumi.set(__self__, "multihop_templates", multihop_templates)
        if neighbors is not None:
            pulumi.set(__self__, "neighbors", neighbors)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="multihopTemplates")
    def multihop_templates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]]]:
        """
        BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        """
        return pulumi.get(self, "multihop_templates")

    @multihop_templates.setter
    def multihop_templates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BfdMultihopTemplateArgs']]]]):
        pulumi.set(self, "multihop_templates", value)

    @property
    @pulumi.getter
    def neighbors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]]]:
        """
        neighbor The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @neighbors.setter
    def neighbors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BfdNeighborArgs']]]]):
        pulumi.set(self, "neighbors", value)

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


class Bfd(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 multihop_templates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdMultihopTemplateArgs']]]]] = None,
                 neighbors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdNeighborArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure BFD.

        ## Import

        Router Bfd can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:router/bfd:Bfd labelname RouterBfd
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:router/bfd:Bfd labelname RouterBfd
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdMultihopTemplateArgs']]]] multihop_templates: BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdNeighborArgs']]]] neighbors: neighbor The structure of `neighbor` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BfdArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure BFD.

        ## Import

        Router Bfd can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:router/bfd:Bfd labelname RouterBfd
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:router/bfd:Bfd labelname RouterBfd
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param BfdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BfdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 multihop_templates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdMultihopTemplateArgs']]]]] = None,
                 neighbors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdNeighborArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BfdArgs.__new__(BfdArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["multihop_templates"] = multihop_templates
            __props__.__dict__["neighbors"] = neighbors
            __props__.__dict__["vdomparam"] = vdomparam
        super(Bfd, __self__).__init__(
            'fortios:router/bfd:Bfd',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            multihop_templates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdMultihopTemplateArgs']]]]] = None,
            neighbors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdNeighborArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Bfd':
        """
        Get an existing Bfd resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdMultihopTemplateArgs']]]] multihop_templates: BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BfdNeighborArgs']]]] neighbors: neighbor The structure of `neighbor` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BfdState.__new__(_BfdState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["multihop_templates"] = multihop_templates
        __props__.__dict__["neighbors"] = neighbors
        __props__.__dict__["vdomparam"] = vdomparam
        return Bfd(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="multihopTemplates")
    def multihop_templates(self) -> pulumi.Output[Optional[Sequence['outputs.BfdMultihopTemplate']]]:
        """
        BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        """
        return pulumi.get(self, "multihop_templates")

    @property
    @pulumi.getter
    def neighbors(self) -> pulumi.Output[Optional[Sequence['outputs.BfdNeighbor']]]:
        """
        neighbor The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
