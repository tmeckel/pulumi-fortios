# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CustomcommandArgs', 'Customcommand']

@pulumi.input_type
class CustomcommandArgs:
    def __init__(__self__, *,
                 command: pulumi.Input[str],
                 command_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Customcommand resource.
        :param pulumi.Input[str] command: String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        :param pulumi.Input[str] command_name: Command name called by the FortiGate switch controller in the execute command.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "command", command)
        if command_name is not None:
            pulumi.set(__self__, "command_name", command_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Input[str]:
        """
        String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: pulumi.Input[str]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter(name="commandName")
    def command_name(self) -> Optional[pulumi.Input[str]]:
        """
        Command name called by the FortiGate switch controller in the execute command.
        """
        return pulumi.get(self, "command_name")

    @command_name.setter
    def command_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command_name", value)

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
class _CustomcommandState:
    def __init__(__self__, *,
                 command: Optional[pulumi.Input[str]] = None,
                 command_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Customcommand resources.
        :param pulumi.Input[str] command: String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        :param pulumi.Input[str] command_name: Command name called by the FortiGate switch controller in the execute command.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if command is not None:
            pulumi.set(__self__, "command", command)
        if command_name is not None:
            pulumi.set(__self__, "command_name", command_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def command(self) -> Optional[pulumi.Input[str]]:
        """
        String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter(name="commandName")
    def command_name(self) -> Optional[pulumi.Input[str]]:
        """
        Command name called by the FortiGate switch controller in the execute command.
        """
        return pulumi.get(self, "command_name")

    @command_name.setter
    def command_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command_name", value)

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
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Customcommand(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command: Optional[pulumi.Input[str]] = None,
                 command_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.switchcontroller.Customcommand("trname",
            command="ls",
            command_name="1")
        ```

        ## Import

        SwitchController CustomCommand can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:switchcontroller/customcommand:Customcommand labelname {{command_name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:switchcontroller/customcommand:Customcommand labelname {{command_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command: String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        :param pulumi.Input[str] command_name: Command name called by the FortiGate switch controller in the execute command.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomcommandArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.switchcontroller.Customcommand("trname",
            command="ls",
            command_name="1")
        ```

        ## Import

        SwitchController CustomCommand can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:switchcontroller/customcommand:Customcommand labelname {{command_name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:switchcontroller/customcommand:Customcommand labelname {{command_name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param CustomcommandArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomcommandArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command: Optional[pulumi.Input[str]] = None,
                 command_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomcommandArgs.__new__(CustomcommandArgs)

            if command is None and not opts.urn:
                raise TypeError("Missing required property 'command'")
            __props__.__dict__["command"] = command
            __props__.__dict__["command_name"] = command_name
            __props__.__dict__["description"] = description
            __props__.__dict__["vdomparam"] = vdomparam
        super(Customcommand, __self__).__init__(
            'fortios:switchcontroller/customcommand:Customcommand',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command: Optional[pulumi.Input[str]] = None,
            command_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Customcommand':
        """
        Get an existing Customcommand resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command: String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        :param pulumi.Input[str] command_name: Command name called by the FortiGate switch controller in the execute command.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomcommandState.__new__(_CustomcommandState)

        __props__.__dict__["command"] = command
        __props__.__dict__["command_name"] = command_name
        __props__.__dict__["description"] = description
        __props__.__dict__["vdomparam"] = vdomparam
        return Customcommand(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output[str]:
        """
        String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter(name="commandName")
    def command_name(self) -> pulumi.Output[str]:
        """
        Command name called by the FortiGate switch controller in the execute command.
        """
        return pulumi.get(self, "command_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
