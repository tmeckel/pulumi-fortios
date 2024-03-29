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

__all__ = ['MpskprofileArgs', 'Mpskprofile']

@pulumi.input_type
class MpskprofileArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Mpskprofile resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] mpsk_concurrent_clients: Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        :param pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]] mpsk_groups: List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        :param pulumi.Input[str] name: MPSK profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mpsk_concurrent_clients is not None:
            pulumi.set(__self__, "mpsk_concurrent_clients", mpsk_concurrent_clients)
        if mpsk_groups is not None:
            pulumi.set(__self__, "mpsk_groups", mpsk_groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="mpskConcurrentClients")
    def mpsk_concurrent_clients(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        """
        return pulumi.get(self, "mpsk_concurrent_clients")

    @mpsk_concurrent_clients.setter
    def mpsk_concurrent_clients(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mpsk_concurrent_clients", value)

    @property
    @pulumi.getter(name="mpskGroups")
    def mpsk_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]]]:
        """
        List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        """
        return pulumi.get(self, "mpsk_groups")

    @mpsk_groups.setter
    def mpsk_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]]]):
        pulumi.set(self, "mpsk_groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        MPSK profile name.
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
class _MpskprofileState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Mpskprofile resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] mpsk_concurrent_clients: Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        :param pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]] mpsk_groups: List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        :param pulumi.Input[str] name: MPSK profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mpsk_concurrent_clients is not None:
            pulumi.set(__self__, "mpsk_concurrent_clients", mpsk_concurrent_clients)
        if mpsk_groups is not None:
            pulumi.set(__self__, "mpsk_groups", mpsk_groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="mpskConcurrentClients")
    def mpsk_concurrent_clients(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        """
        return pulumi.get(self, "mpsk_concurrent_clients")

    @mpsk_concurrent_clients.setter
    def mpsk_concurrent_clients(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mpsk_concurrent_clients", value)

    @property
    @pulumi.getter(name="mpskGroups")
    def mpsk_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]]]:
        """
        List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        """
        return pulumi.get(self, "mpsk_groups")

    @mpsk_groups.setter
    def mpsk_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MpskprofileMpskGroupArgs']]]]):
        pulumi.set(self, "mpsk_groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        MPSK profile name.
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


class Mpskprofile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MpskprofileMpskGroupArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure MPSK profile. Applies to FortiOS Version `>= 6.4.2`.

        ## Import

        WirelessController MpskProfile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:wirelesscontroller/mpskprofile:Mpskprofile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:wirelesscontroller/mpskprofile:Mpskprofile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] mpsk_concurrent_clients: Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MpskprofileMpskGroupArgs']]]] mpsk_groups: List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        :param pulumi.Input[str] name: MPSK profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MpskprofileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure MPSK profile. Applies to FortiOS Version `>= 6.4.2`.

        ## Import

        WirelessController MpskProfile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:wirelesscontroller/mpskprofile:Mpskprofile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:wirelesscontroller/mpskprofile:Mpskprofile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param MpskprofileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MpskprofileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MpskprofileMpskGroupArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MpskprofileArgs.__new__(MpskprofileArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["mpsk_concurrent_clients"] = mpsk_concurrent_clients
            __props__.__dict__["mpsk_groups"] = mpsk_groups
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Mpskprofile, __self__).__init__(
            'fortios:wirelesscontroller/mpskprofile:Mpskprofile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
            mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MpskprofileMpskGroupArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Mpskprofile':
        """
        Get an existing Mpskprofile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] mpsk_concurrent_clients: Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MpskprofileMpskGroupArgs']]]] mpsk_groups: List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        :param pulumi.Input[str] name: MPSK profile name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MpskprofileState.__new__(_MpskprofileState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["mpsk_concurrent_clients"] = mpsk_concurrent_clients
        __props__.__dict__["mpsk_groups"] = mpsk_groups
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Mpskprofile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="mpskConcurrentClients")
    def mpsk_concurrent_clients(self) -> pulumi.Output[int]:
        """
        Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
        """
        return pulumi.get(self, "mpsk_concurrent_clients")

    @property
    @pulumi.getter(name="mpskGroups")
    def mpsk_groups(self) -> pulumi.Output[Optional[Sequence['outputs.MpskprofileMpskGroup']]]:
        """
        List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
        """
        return pulumi.get(self, "mpsk_groups")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        MPSK profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

