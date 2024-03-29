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

__all__ = ['LocationArgs', 'Location']

@pulumi.input_type
class LocationArgs:
    def __init__(__self__, *,
                 address_civic: Optional[pulumi.Input['LocationAddressCivicArgs']] = None,
                 coordinates: Optional[pulumi.Input['LocationCoordinatesArgs']] = None,
                 elin_number: Optional[pulumi.Input['LocationElinNumberArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Location resource.
        :param pulumi.Input['LocationAddressCivicArgs'] address_civic: Configure location civic address. The structure of `address_civic` block is documented below.
        :param pulumi.Input['LocationCoordinatesArgs'] coordinates: Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        :param pulumi.Input['LocationElinNumberArgs'] elin_number: Configure location ELIN number. The structure of `elin_number` block is documented below.
        :param pulumi.Input[str] name: Unique location item name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if address_civic is not None:
            pulumi.set(__self__, "address_civic", address_civic)
        if coordinates is not None:
            pulumi.set(__self__, "coordinates", coordinates)
        if elin_number is not None:
            pulumi.set(__self__, "elin_number", elin_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addressCivic")
    def address_civic(self) -> Optional[pulumi.Input['LocationAddressCivicArgs']]:
        """
        Configure location civic address. The structure of `address_civic` block is documented below.
        """
        return pulumi.get(self, "address_civic")

    @address_civic.setter
    def address_civic(self, value: Optional[pulumi.Input['LocationAddressCivicArgs']]):
        pulumi.set(self, "address_civic", value)

    @property
    @pulumi.getter
    def coordinates(self) -> Optional[pulumi.Input['LocationCoordinatesArgs']]:
        """
        Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        """
        return pulumi.get(self, "coordinates")

    @coordinates.setter
    def coordinates(self, value: Optional[pulumi.Input['LocationCoordinatesArgs']]):
        pulumi.set(self, "coordinates", value)

    @property
    @pulumi.getter(name="elinNumber")
    def elin_number(self) -> Optional[pulumi.Input['LocationElinNumberArgs']]:
        """
        Configure location ELIN number. The structure of `elin_number` block is documented below.
        """
        return pulumi.get(self, "elin_number")

    @elin_number.setter
    def elin_number(self, value: Optional[pulumi.Input['LocationElinNumberArgs']]):
        pulumi.set(self, "elin_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique location item name.
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
class _LocationState:
    def __init__(__self__, *,
                 address_civic: Optional[pulumi.Input['LocationAddressCivicArgs']] = None,
                 coordinates: Optional[pulumi.Input['LocationCoordinatesArgs']] = None,
                 elin_number: Optional[pulumi.Input['LocationElinNumberArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Location resources.
        :param pulumi.Input['LocationAddressCivicArgs'] address_civic: Configure location civic address. The structure of `address_civic` block is documented below.
        :param pulumi.Input['LocationCoordinatesArgs'] coordinates: Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        :param pulumi.Input['LocationElinNumberArgs'] elin_number: Configure location ELIN number. The structure of `elin_number` block is documented below.
        :param pulumi.Input[str] name: Unique location item name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if address_civic is not None:
            pulumi.set(__self__, "address_civic", address_civic)
        if coordinates is not None:
            pulumi.set(__self__, "coordinates", coordinates)
        if elin_number is not None:
            pulumi.set(__self__, "elin_number", elin_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addressCivic")
    def address_civic(self) -> Optional[pulumi.Input['LocationAddressCivicArgs']]:
        """
        Configure location civic address. The structure of `address_civic` block is documented below.
        """
        return pulumi.get(self, "address_civic")

    @address_civic.setter
    def address_civic(self, value: Optional[pulumi.Input['LocationAddressCivicArgs']]):
        pulumi.set(self, "address_civic", value)

    @property
    @pulumi.getter
    def coordinates(self) -> Optional[pulumi.Input['LocationCoordinatesArgs']]:
        """
        Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        """
        return pulumi.get(self, "coordinates")

    @coordinates.setter
    def coordinates(self, value: Optional[pulumi.Input['LocationCoordinatesArgs']]):
        pulumi.set(self, "coordinates", value)

    @property
    @pulumi.getter(name="elinNumber")
    def elin_number(self) -> Optional[pulumi.Input['LocationElinNumberArgs']]:
        """
        Configure location ELIN number. The structure of `elin_number` block is documented below.
        """
        return pulumi.get(self, "elin_number")

    @elin_number.setter
    def elin_number(self, value: Optional[pulumi.Input['LocationElinNumberArgs']]):
        pulumi.set(self, "elin_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique location item name.
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


class Location(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_civic: Optional[pulumi.Input[pulumi.InputType['LocationAddressCivicArgs']]] = None,
                 coordinates: Optional[pulumi.Input[pulumi.InputType['LocationCoordinatesArgs']]] = None,
                 elin_number: Optional[pulumi.Input[pulumi.InputType['LocationElinNumberArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSwitch location services. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController Location can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LocationAddressCivicArgs']] address_civic: Configure location civic address. The structure of `address_civic` block is documented below.
        :param pulumi.Input[pulumi.InputType['LocationCoordinatesArgs']] coordinates: Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        :param pulumi.Input[pulumi.InputType['LocationElinNumberArgs']] elin_number: Configure location ELIN number. The structure of `elin_number` block is documented below.
        :param pulumi.Input[str] name: Unique location item name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LocationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSwitch location services. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController Location can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:switchcontroller/location:Location labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param LocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_civic: Optional[pulumi.Input[pulumi.InputType['LocationAddressCivicArgs']]] = None,
                 coordinates: Optional[pulumi.Input[pulumi.InputType['LocationCoordinatesArgs']]] = None,
                 elin_number: Optional[pulumi.Input[pulumi.InputType['LocationElinNumberArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocationArgs.__new__(LocationArgs)

            __props__.__dict__["address_civic"] = address_civic
            __props__.__dict__["coordinates"] = coordinates
            __props__.__dict__["elin_number"] = elin_number
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Location, __self__).__init__(
            'fortios:switchcontroller/location:Location',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_civic: Optional[pulumi.Input[pulumi.InputType['LocationAddressCivicArgs']]] = None,
            coordinates: Optional[pulumi.Input[pulumi.InputType['LocationCoordinatesArgs']]] = None,
            elin_number: Optional[pulumi.Input[pulumi.InputType['LocationElinNumberArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Location':
        """
        Get an existing Location resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LocationAddressCivicArgs']] address_civic: Configure location civic address. The structure of `address_civic` block is documented below.
        :param pulumi.Input[pulumi.InputType['LocationCoordinatesArgs']] coordinates: Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        :param pulumi.Input[pulumi.InputType['LocationElinNumberArgs']] elin_number: Configure location ELIN number. The structure of `elin_number` block is documented below.
        :param pulumi.Input[str] name: Unique location item name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LocationState.__new__(_LocationState)

        __props__.__dict__["address_civic"] = address_civic
        __props__.__dict__["coordinates"] = coordinates
        __props__.__dict__["elin_number"] = elin_number
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Location(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressCivic")
    def address_civic(self) -> pulumi.Output['outputs.LocationAddressCivic']:
        """
        Configure location civic address. The structure of `address_civic` block is documented below.
        """
        return pulumi.get(self, "address_civic")

    @property
    @pulumi.getter
    def coordinates(self) -> pulumi.Output['outputs.LocationCoordinates']:
        """
        Configure location GPS coordinates. The structure of `coordinates` block is documented below.
        """
        return pulumi.get(self, "coordinates")

    @property
    @pulumi.getter(name="elinNumber")
    def elin_number(self) -> pulumi.Output['outputs.LocationElinNumber']:
        """
        Configure location ELIN number. The structure of `elin_number` block is documented below.
        """
        return pulumi.get(self, "elin_number")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique location item name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

