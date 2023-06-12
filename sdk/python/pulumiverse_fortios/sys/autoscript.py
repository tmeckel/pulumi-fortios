# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AutoscriptArgs', 'Autoscript']

@pulumi.input_type
class AutoscriptArgs:
    def __init__(__self__, *,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_size: Optional[pulumi.Input[int]] = None,
                 repeat: Optional[pulumi.Input[int]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Autoscript resource.
        :param pulumi.Input[int] interval: Repeat interval in seconds.
        :param pulumi.Input[str] name: Auto script name.
        :param pulumi.Input[int] output_size: Number of megabytes to limit script output to (10 - 1024, default = 10).
        :param pulumi.Input[int] repeat: Number of times to repeat this script (0 = infinite).
        :param pulumi.Input[str] script: List of FortiOS CLI commands to repeat.
        :param pulumi.Input[str] start: Script starting mode. Valid values: `manual`, `auto`.
        :param pulumi.Input[int] timeout: Maximum running time for this script in seconds (0 = no timeout).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if output_size is not None:
            pulumi.set(__self__, "output_size", output_size)
        if repeat is not None:
            pulumi.set(__self__, "repeat", repeat)
        if script is not None:
            pulumi.set(__self__, "script", script)
        if start is not None:
            pulumi.set(__self__, "start", start)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Repeat interval in seconds.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Auto script name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outputSize")
    def output_size(self) -> Optional[pulumi.Input[int]]:
        """
        Number of megabytes to limit script output to (10 - 1024, default = 10).
        """
        return pulumi.get(self, "output_size")

    @output_size.setter
    def output_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "output_size", value)

    @property
    @pulumi.getter
    def repeat(self) -> Optional[pulumi.Input[int]]:
        """
        Number of times to repeat this script (0 = infinite).
        """
        return pulumi.get(self, "repeat")

    @repeat.setter
    def repeat(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat", value)

    @property
    @pulumi.getter
    def script(self) -> Optional[pulumi.Input[str]]:
        """
        List of FortiOS CLI commands to repeat.
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[str]]:
        """
        Script starting mode. Valid values: `manual`, `auto`.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum running time for this script in seconds (0 = no timeout).
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

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
class _AutoscriptState:
    def __init__(__self__, *,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_size: Optional[pulumi.Input[int]] = None,
                 repeat: Optional[pulumi.Input[int]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Autoscript resources.
        :param pulumi.Input[int] interval: Repeat interval in seconds.
        :param pulumi.Input[str] name: Auto script name.
        :param pulumi.Input[int] output_size: Number of megabytes to limit script output to (10 - 1024, default = 10).
        :param pulumi.Input[int] repeat: Number of times to repeat this script (0 = infinite).
        :param pulumi.Input[str] script: List of FortiOS CLI commands to repeat.
        :param pulumi.Input[str] start: Script starting mode. Valid values: `manual`, `auto`.
        :param pulumi.Input[int] timeout: Maximum running time for this script in seconds (0 = no timeout).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if output_size is not None:
            pulumi.set(__self__, "output_size", output_size)
        if repeat is not None:
            pulumi.set(__self__, "repeat", repeat)
        if script is not None:
            pulumi.set(__self__, "script", script)
        if start is not None:
            pulumi.set(__self__, "start", start)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Repeat interval in seconds.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Auto script name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outputSize")
    def output_size(self) -> Optional[pulumi.Input[int]]:
        """
        Number of megabytes to limit script output to (10 - 1024, default = 10).
        """
        return pulumi.get(self, "output_size")

    @output_size.setter
    def output_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "output_size", value)

    @property
    @pulumi.getter
    def repeat(self) -> Optional[pulumi.Input[int]]:
        """
        Number of times to repeat this script (0 = infinite).
        """
        return pulumi.get(self, "repeat")

    @repeat.setter
    def repeat(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat", value)

    @property
    @pulumi.getter
    def script(self) -> Optional[pulumi.Input[str]]:
        """
        List of FortiOS CLI commands to repeat.
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[str]]:
        """
        Script starting mode. Valid values: `manual`, `auto`.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum running time for this script in seconds (0 = no timeout).
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

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


class Autoscript(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_size: Optional[pulumi.Input[int]] = None,
                 repeat: Optional[pulumi.Input[int]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure auto script.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        auto2 = fortios.sys.Autoscript("auto2",
            interval=1,
            output_size=10,
            repeat=1,
            script=\"\"\"config firewall address
            edit "111"
                set color 3
                set subnet 1.1.1.1 255.255.255.255
            next
        end

        \"\"\",
            start="auto")
        ```

        ## Import

        System AutoScript can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/autoscript:Autoscript labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/autoscript:Autoscript labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] interval: Repeat interval in seconds.
        :param pulumi.Input[str] name: Auto script name.
        :param pulumi.Input[int] output_size: Number of megabytes to limit script output to (10 - 1024, default = 10).
        :param pulumi.Input[int] repeat: Number of times to repeat this script (0 = infinite).
        :param pulumi.Input[str] script: List of FortiOS CLI commands to repeat.
        :param pulumi.Input[str] start: Script starting mode. Valid values: `manual`, `auto`.
        :param pulumi.Input[int] timeout: Maximum running time for this script in seconds (0 = no timeout).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AutoscriptArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure auto script.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        auto2 = fortios.sys.Autoscript("auto2",
            interval=1,
            output_size=10,
            repeat=1,
            script=\"\"\"config firewall address
            edit "111"
                set color 3
                set subnet 1.1.1.1 255.255.255.255
            next
        end

        \"\"\",
            start="auto")
        ```

        ## Import

        System AutoScript can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:sys/autoscript:Autoscript labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:sys/autoscript:Autoscript labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param AutoscriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoscriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_size: Optional[pulumi.Input[int]] = None,
                 repeat: Optional[pulumi.Input[int]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoscriptArgs.__new__(AutoscriptArgs)

            __props__.__dict__["interval"] = interval
            __props__.__dict__["name"] = name
            __props__.__dict__["output_size"] = output_size
            __props__.__dict__["repeat"] = repeat
            __props__.__dict__["script"] = script
            __props__.__dict__["start"] = start
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["vdomparam"] = vdomparam
        super(Autoscript, __self__).__init__(
            'fortios:sys/autoscript:Autoscript',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            interval: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            output_size: Optional[pulumi.Input[int]] = None,
            repeat: Optional[pulumi.Input[int]] = None,
            script: Optional[pulumi.Input[str]] = None,
            start: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Autoscript':
        """
        Get an existing Autoscript resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] interval: Repeat interval in seconds.
        :param pulumi.Input[str] name: Auto script name.
        :param pulumi.Input[int] output_size: Number of megabytes to limit script output to (10 - 1024, default = 10).
        :param pulumi.Input[int] repeat: Number of times to repeat this script (0 = infinite).
        :param pulumi.Input[str] script: List of FortiOS CLI commands to repeat.
        :param pulumi.Input[str] start: Script starting mode. Valid values: `manual`, `auto`.
        :param pulumi.Input[int] timeout: Maximum running time for this script in seconds (0 = no timeout).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutoscriptState.__new__(_AutoscriptState)

        __props__.__dict__["interval"] = interval
        __props__.__dict__["name"] = name
        __props__.__dict__["output_size"] = output_size
        __props__.__dict__["repeat"] = repeat
        __props__.__dict__["script"] = script
        __props__.__dict__["start"] = start
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["vdomparam"] = vdomparam
        return Autoscript(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[int]:
        """
        Repeat interval in seconds.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Auto script name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputSize")
    def output_size(self) -> pulumi.Output[int]:
        """
        Number of megabytes to limit script output to (10 - 1024, default = 10).
        """
        return pulumi.get(self, "output_size")

    @property
    @pulumi.getter
    def repeat(self) -> pulumi.Output[int]:
        """
        Number of times to repeat this script (0 = infinite).
        """
        return pulumi.get(self, "repeat")

    @property
    @pulumi.getter
    def script(self) -> pulumi.Output[Optional[str]]:
        """
        List of FortiOS CLI commands to repeat.
        """
        return pulumi.get(self, "script")

    @property
    @pulumi.getter
    def start(self) -> pulumi.Output[str]:
        """
        Script starting mode. Valid values: `manual`, `auto`.
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        Maximum running time for this script in seconds (0 = no timeout).
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

