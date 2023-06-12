# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SystemNetworkRouteArgs', 'SystemNetworkRoute']

@pulumi.input_type
class SystemNetworkRouteArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str],
                 device: pulumi.Input[str],
                 gateway: pulumi.Input[str],
                 route_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a SystemNetworkRoute resource.
        :param pulumi.Input[str] destination: Destination Ip/Mask.
        :param pulumi.Input[str] device: Gateway out interface.
        :param pulumi.Input[str] gateway: Gateway Ip.
        :param pulumi.Input[int] route_id: Route id.
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "device", device)
        pulumi.set(__self__, "gateway", gateway)
        pulumi.set(__self__, "route_id", route_id)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        Destination Ip/Mask.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Input[str]:
        """
        Gateway out interface.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: pulumi.Input[str]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Input[str]:
        """
        Gateway Ip.
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Input[int]:
        """
        Route id.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "route_id", value)


@pulumi.input_type
class _SystemNetworkRouteState:
    def __init__(__self__, *,
                 destination: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SystemNetworkRoute resources.
        :param pulumi.Input[str] destination: Destination Ip/Mask.
        :param pulumi.Input[str] device: Gateway out interface.
        :param pulumi.Input[str] gateway: Gateway Ip.
        :param pulumi.Input[int] route_id: Route id.
        """
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if route_id is not None:
            pulumi.set(__self__, "route_id", route_id)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        """
        Destination Ip/Mask.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway out interface.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway Ip.
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> Optional[pulumi.Input[int]]:
        """
        Route id.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "route_id", value)


class SystemNetworkRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource supports updating system network route for FortiManager.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.SystemNetworkRoute("test1",
            destination="192.168.2.0 255.255.255.0",
            device="port4",
            gateway="192.168.2.1",
            route_id=5)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: Destination Ip/Mask.
        :param pulumi.Input[str] device: Gateway out interface.
        :param pulumi.Input[str] gateway: Gateway Ip.
        :param pulumi.Input[int] route_id: Route id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemNetworkRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports updating system network route for FortiManager.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test1 = fortios.fmg.SystemNetworkRoute("test1",
            destination="192.168.2.0 255.255.255.0",
            device="port4",
            gateway="192.168.2.1",
            route_id=5)
        ```

        :param str resource_name: The name of the resource.
        :param SystemNetworkRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemNetworkRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemNetworkRouteArgs.__new__(SystemNetworkRouteArgs)

            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__["destination"] = destination
            if device is None and not opts.urn:
                raise TypeError("Missing required property 'device'")
            __props__.__dict__["device"] = device
            if gateway is None and not opts.urn:
                raise TypeError("Missing required property 'gateway'")
            __props__.__dict__["gateway"] = gateway
            if route_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_id'")
            __props__.__dict__["route_id"] = route_id
        super(SystemNetworkRoute, __self__).__init__(
            'fortios:fmg/systemNetworkRoute:SystemNetworkRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            route_id: Optional[pulumi.Input[int]] = None) -> 'SystemNetworkRoute':
        """
        Get an existing SystemNetworkRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: Destination Ip/Mask.
        :param pulumi.Input[str] device: Gateway out interface.
        :param pulumi.Input[str] gateway: Gateway Ip.
        :param pulumi.Input[int] route_id: Route id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemNetworkRouteState.__new__(_SystemNetworkRouteState)

        __props__.__dict__["destination"] = destination
        __props__.__dict__["device"] = device
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["route_id"] = route_id
        return SystemNetworkRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        Destination Ip/Mask.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        """
        Gateway out interface.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        """
        Gateway Ip.
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Output[int]:
        """
        Route id.
        """
        return pulumi.get(self, "route_id")
