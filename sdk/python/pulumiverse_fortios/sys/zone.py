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

__all__ = ['ZoneArgs', 'Zone']

@pulumi.input_type
class ZoneArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]]] = None,
                 intrazone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Zone resource.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]] interfaces: Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        :param pulumi.Input[str] intrazone: Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        :param pulumi.Input[str] name: Zone name.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if intrazone is not None:
            pulumi.set(__self__, "intrazone", intrazone)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]]]:
        """
        Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter
    def intrazone(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "intrazone")

    @intrazone.setter
    def intrazone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intrazone", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Zone name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]]]:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]]]):
        pulumi.set(self, "taggings", value)

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
class _ZoneState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]]] = None,
                 intrazone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Zone resources.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]] interfaces: Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        :param pulumi.Input[str] intrazone: Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        :param pulumi.Input[str] name: Zone name.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if intrazone is not None:
            pulumi.set(__self__, "intrazone", intrazone)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]]]:
        """
        Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter
    def intrazone(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "intrazone")

    @intrazone.setter
    def intrazone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intrazone", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Zone name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]]]:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTaggingArgs']]]]):
        pulumi.set(self, "taggings", value)

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


class Zone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneInterfaceArgs']]]]] = None,
                 intrazone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTaggingArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Zone("trname", intrazone="allow")
        ```

        ## Import

        System Zone can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/zone:Zone labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/zone:Zone labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneInterfaceArgs']]]] interfaces: Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        :param pulumi.Input[str] intrazone: Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        :param pulumi.Input[str] name: Zone name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTaggingArgs']]]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ZoneArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.sys.Zone("trname", intrazone="allow")
        ```

        ## Import

        System Zone can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/zone:Zone labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/zone:Zone labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneInterfaceArgs']]]]] = None,
                 intrazone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTaggingArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneArgs.__new__(ZoneArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["intrazone"] = intrazone
            __props__.__dict__["name"] = name
            __props__.__dict__["taggings"] = taggings
            __props__.__dict__["vdomparam"] = vdomparam
        super(Zone, __self__).__init__(
            'fortios:sys/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneInterfaceArgs']]]]] = None,
            intrazone: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTaggingArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneInterfaceArgs']]]] interfaces: Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        :param pulumi.Input[str] intrazone: Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        :param pulumi.Input[str] name: Zone name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTaggingArgs']]]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneState.__new__(_ZoneState)

        __props__.__dict__["description"] = description
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["intrazone"] = intrazone
        __props__.__dict__["name"] = name
        __props__.__dict__["taggings"] = taggings
        __props__.__dict__["vdomparam"] = vdomparam
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.ZoneInterface']]]:
        """
        Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    def intrazone(self) -> pulumi.Output[str]:
        """
        Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "intrazone")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Zone name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> pulumi.Output[Optional[Sequence['outputs.ZoneTagging']]]:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
