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

__all__ = ['QosmapArgs', 'Qosmap']

@pulumi.input_type
class QosmapArgs:
    def __init__(__self__, *,
                 dscp_excepts: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]]] = None,
                 dscp_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Qosmap resource.
        :param pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]] dscp_excepts: Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]] dscp_ranges: Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: QOS-MAP name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dscp_excepts is not None:
            pulumi.set(__self__, "dscp_excepts", dscp_excepts)
        if dscp_ranges is not None:
            pulumi.set(__self__, "dscp_ranges", dscp_ranges)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dscpExcepts")
    def dscp_excepts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]]]:
        """
        Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        """
        return pulumi.get(self, "dscp_excepts")

    @dscp_excepts.setter
    def dscp_excepts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]]]):
        pulumi.set(self, "dscp_excepts", value)

    @property
    @pulumi.getter(name="dscpRanges")
    def dscp_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]]]:
        """
        Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        """
        return pulumi.get(self, "dscp_ranges")

    @dscp_ranges.setter
    def dscp_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]]]):
        pulumi.set(self, "dscp_ranges", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        QOS-MAP name.
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
class _QosmapState:
    def __init__(__self__, *,
                 dscp_excepts: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]]] = None,
                 dscp_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Qosmap resources.
        :param pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]] dscp_excepts: Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]] dscp_ranges: Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: QOS-MAP name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dscp_excepts is not None:
            pulumi.set(__self__, "dscp_excepts", dscp_excepts)
        if dscp_ranges is not None:
            pulumi.set(__self__, "dscp_ranges", dscp_ranges)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dscpExcepts")
    def dscp_excepts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]]]:
        """
        Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        """
        return pulumi.get(self, "dscp_excepts")

    @dscp_excepts.setter
    def dscp_excepts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpExceptArgs']]]]):
        pulumi.set(self, "dscp_excepts", value)

    @property
    @pulumi.getter(name="dscpRanges")
    def dscp_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]]]:
        """
        Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        """
        return pulumi.get(self, "dscp_ranges")

    @dscp_ranges.setter
    def dscp_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['QosmapDscpRangeArgs']]]]):
        pulumi.set(self, "dscp_ranges", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        QOS-MAP name.
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


class Qosmap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dscp_excepts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpExceptArgs']]]]] = None,
                 dscp_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpRangeArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure QoS map set.

        ## Import

        WirelessControllerHotspot20 QosMap can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:wirelesscontrollerhotspot20/qosmap:Qosmap labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:wirelesscontrollerhotspot20/qosmap:Qosmap labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpExceptArgs']]]] dscp_excepts: Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpRangeArgs']]]] dscp_ranges: Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: QOS-MAP name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[QosmapArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure QoS map set.

        ## Import

        WirelessControllerHotspot20 QosMap can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:wirelesscontrollerhotspot20/qosmap:Qosmap labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:wirelesscontrollerhotspot20/qosmap:Qosmap labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param QosmapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QosmapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dscp_excepts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpExceptArgs']]]]] = None,
                 dscp_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpRangeArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QosmapArgs.__new__(QosmapArgs)

            __props__.__dict__["dscp_excepts"] = dscp_excepts
            __props__.__dict__["dscp_ranges"] = dscp_ranges
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Qosmap, __self__).__init__(
            'fortios:wirelesscontrollerhotspot20/qosmap:Qosmap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dscp_excepts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpExceptArgs']]]]] = None,
            dscp_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpRangeArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Qosmap':
        """
        Get an existing Qosmap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpExceptArgs']]]] dscp_excepts: Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['QosmapDscpRangeArgs']]]] dscp_ranges: Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: QOS-MAP name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QosmapState.__new__(_QosmapState)

        __props__.__dict__["dscp_excepts"] = dscp_excepts
        __props__.__dict__["dscp_ranges"] = dscp_ranges
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Qosmap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dscpExcepts")
    def dscp_excepts(self) -> pulumi.Output[Optional[Sequence['outputs.QosmapDscpExcept']]]:
        """
        Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
        """
        return pulumi.get(self, "dscp_excepts")

    @property
    @pulumi.getter(name="dscpRanges")
    def dscp_ranges(self) -> pulumi.Output[Optional[Sequence['outputs.QosmapDscpRange']]]:
        """
        Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.
        """
        return pulumi.get(self, "dscp_ranges")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        QOS-MAP name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

