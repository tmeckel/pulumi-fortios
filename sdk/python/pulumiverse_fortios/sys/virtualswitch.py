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

__all__ = ['VirtualswitchArgs', 'Virtualswitch']

@pulumi.input_type
class VirtualswitchArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_switch: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Virtualswitch resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Name of the virtual switch.
        :param pulumi.Input[str] physical_switch: Physical switch parent.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]] ports: Configure member ports. The structure of `port` block is documented below.
        :param pulumi.Input[str] span: Enable/disable SPAN. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port.
        :param pulumi.Input[str] span_direction: SPAN direction. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[str] span_source_port: SPAN source port.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlan: VLAN.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if physical_switch is not None:
            pulumi.set(__self__, "physical_switch", physical_switch)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if span is not None:
            pulumi.set(__self__, "span", span)
        if span_dest_port is not None:
            pulumi.set(__self__, "span_dest_port", span_dest_port)
        if span_direction is not None:
            pulumi.set(__self__, "span_direction", span_direction)
        if span_source_port is not None:
            pulumi.set(__self__, "span_source_port", span_source_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

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
        Name of the virtual switch.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="physicalSwitch")
    def physical_switch(self) -> Optional[pulumi.Input[str]]:
        """
        Physical switch parent.
        """
        return pulumi.get(self, "physical_switch")

    @physical_switch.setter
    def physical_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "physical_switch", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]]]:
        """
        Configure member ports. The structure of `port` block is documented below.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def span(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable SPAN. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "span")

    @span.setter
    def span(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span", value)

    @property
    @pulumi.getter(name="spanDestPort")
    def span_dest_port(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN destination port.
        """
        return pulumi.get(self, "span_dest_port")

    @span_dest_port.setter
    def span_dest_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_dest_port", value)

    @property
    @pulumi.getter(name="spanDirection")
    def span_direction(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN direction. Valid values: `rx`, `tx`, `both`.
        """
        return pulumi.get(self, "span_direction")

    @span_direction.setter
    def span_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_direction", value)

    @property
    @pulumi.getter(name="spanSourcePort")
    def span_source_port(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN source port.
        """
        return pulumi.get(self, "span_source_port")

    @span_source_port.setter
    def span_source_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_source_port", value)

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
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)


@pulumi.input_type
class _VirtualswitchState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_switch: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Virtualswitch resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Name of the virtual switch.
        :param pulumi.Input[str] physical_switch: Physical switch parent.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]] ports: Configure member ports. The structure of `port` block is documented below.
        :param pulumi.Input[str] span: Enable/disable SPAN. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port.
        :param pulumi.Input[str] span_direction: SPAN direction. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[str] span_source_port: SPAN source port.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlan: VLAN.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if physical_switch is not None:
            pulumi.set(__self__, "physical_switch", physical_switch)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if span is not None:
            pulumi.set(__self__, "span", span)
        if span_dest_port is not None:
            pulumi.set(__self__, "span_dest_port", span_dest_port)
        if span_direction is not None:
            pulumi.set(__self__, "span_direction", span_direction)
        if span_source_port is not None:
            pulumi.set(__self__, "span_source_port", span_source_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

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
        Name of the virtual switch.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="physicalSwitch")
    def physical_switch(self) -> Optional[pulumi.Input[str]]:
        """
        Physical switch parent.
        """
        return pulumi.get(self, "physical_switch")

    @physical_switch.setter
    def physical_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "physical_switch", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]]]:
        """
        Configure member ports. The structure of `port` block is documented below.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualswitchPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def span(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable SPAN. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "span")

    @span.setter
    def span(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span", value)

    @property
    @pulumi.getter(name="spanDestPort")
    def span_dest_port(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN destination port.
        """
        return pulumi.get(self, "span_dest_port")

    @span_dest_port.setter
    def span_dest_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_dest_port", value)

    @property
    @pulumi.getter(name="spanDirection")
    def span_direction(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN direction. Valid values: `rx`, `tx`, `both`.
        """
        return pulumi.get(self, "span_direction")

    @span_direction.setter
    def span_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_direction", value)

    @property
    @pulumi.getter(name="spanSourcePort")
    def span_source_port(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN source port.
        """
        return pulumi.get(self, "span_source_port")

    @span_source_port.setter
    def span_source_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_source_port", value)

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
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)


class Virtualswitch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_switch: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualswitchPortArgs']]]]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configure virtual hardware switch interfaces. Applies to FortiOS Version `7.0.4`.

        ## Import

        System VirtualSwitch can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/virtualswitch:Virtualswitch labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/virtualswitch:Virtualswitch labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Name of the virtual switch.
        :param pulumi.Input[str] physical_switch: Physical switch parent.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualswitchPortArgs']]]] ports: Configure member ports. The structure of `port` block is documented below.
        :param pulumi.Input[str] span: Enable/disable SPAN. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port.
        :param pulumi.Input[str] span_direction: SPAN direction. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[str] span_source_port: SPAN source port.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlan: VLAN.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VirtualswitchArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure virtual hardware switch interfaces. Applies to FortiOS Version `7.0.4`.

        ## Import

        System VirtualSwitch can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/virtualswitch:Virtualswitch labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/virtualswitch:Virtualswitch labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VirtualswitchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualswitchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_switch: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualswitchPortArgs']]]]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualswitchArgs.__new__(VirtualswitchArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["physical_switch"] = physical_switch
            __props__.__dict__["ports"] = ports
            __props__.__dict__["span"] = span
            __props__.__dict__["span_dest_port"] = span_dest_port
            __props__.__dict__["span_direction"] = span_direction
            __props__.__dict__["span_source_port"] = span_source_port
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vlan"] = vlan
        super(Virtualswitch, __self__).__init__(
            'fortios:sys/virtualswitch:Virtualswitch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            physical_switch: Optional[pulumi.Input[str]] = None,
            ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualswitchPortArgs']]]]] = None,
            span: Optional[pulumi.Input[str]] = None,
            span_dest_port: Optional[pulumi.Input[str]] = None,
            span_direction: Optional[pulumi.Input[str]] = None,
            span_source_port: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vlan: Optional[pulumi.Input[int]] = None) -> 'Virtualswitch':
        """
        Get an existing Virtualswitch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Name of the virtual switch.
        :param pulumi.Input[str] physical_switch: Physical switch parent.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualswitchPortArgs']]]] ports: Configure member ports. The structure of `port` block is documented below.
        :param pulumi.Input[str] span: Enable/disable SPAN. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port.
        :param pulumi.Input[str] span_direction: SPAN direction. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[str] span_source_port: SPAN source port.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlan: VLAN.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualswitchState.__new__(_VirtualswitchState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["physical_switch"] = physical_switch
        __props__.__dict__["ports"] = ports
        __props__.__dict__["span"] = span
        __props__.__dict__["span_dest_port"] = span_dest_port
        __props__.__dict__["span_direction"] = span_direction
        __props__.__dict__["span_source_port"] = span_source_port
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vlan"] = vlan
        return Virtualswitch(resource_name, opts=opts, __props__=__props__)

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
        Name of the virtual switch.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="physicalSwitch")
    def physical_switch(self) -> pulumi.Output[str]:
        """
        Physical switch parent.
        """
        return pulumi.get(self, "physical_switch")

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualswitchPort']]]:
        """
        Configure member ports. The structure of `port` block is documented below.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def span(self) -> pulumi.Output[str]:
        """
        Enable/disable SPAN. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "span")

    @property
    @pulumi.getter(name="spanDestPort")
    def span_dest_port(self) -> pulumi.Output[str]:
        """
        SPAN destination port.
        """
        return pulumi.get(self, "span_dest_port")

    @property
    @pulumi.getter(name="spanDirection")
    def span_direction(self) -> pulumi.Output[str]:
        """
        SPAN direction. Valid values: `rx`, `tx`, `both`.
        """
        return pulumi.get(self, "span_direction")

    @property
    @pulumi.getter(name="spanSourcePort")
    def span_source_port(self) -> pulumi.Output[str]:
        """
        SPAN source port.
        """
        return pulumi.get(self, "span_source_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[int]:
        """
        VLAN.
        """
        return pulumi.get(self, "vlan")

