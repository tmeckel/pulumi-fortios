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

__all__ = ['Extender1Args', 'Extender1']

@pulumi.input_type
class Extender1Args:
    def __init__(__self__, *,
                 authorized: pulumi.Input[str],
                 controller_report: Optional[pulumi.Input['Extender1ControllerReportArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_name: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 login_password: Optional[pulumi.Input[str]] = None,
                 modem1: Optional[pulumi.Input['Extender1Modem1Args']] = None,
                 modem2: Optional[pulumi.Input['Extender1Modem2Args']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Extender1 resource.
        :param pulumi.Input[str] authorized: FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        :param pulumi.Input['Extender1ControllerReportArgs'] controller_report: FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] ext_name: FortiExtender name.
        :param pulumi.Input[str] fosid: FortiExtender serial number.
        :param pulumi.Input[str] login_password: FortiExtender login password.
        :param pulumi.Input['Extender1Modem1Args'] modem1: Configuration options for modem 1. The structure of `modem1` block is documented below.
        :param pulumi.Input['Extender1Modem2Args'] modem2: Configuration options for modem 2. The structure of `modem2` block is documented below.
        :param pulumi.Input[str] name: FortiExtender entry name.
        :param pulumi.Input[int] vdom: VDOM
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "authorized", authorized)
        if controller_report is not None:
            pulumi.set(__self__, "controller_report", controller_report)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ext_name is not None:
            pulumi.set(__self__, "ext_name", ext_name)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if login_password is not None:
            pulumi.set(__self__, "login_password", login_password)
        if modem1 is not None:
            pulumi.set(__self__, "modem1", modem1)
        if modem2 is not None:
            pulumi.set(__self__, "modem2", modem2)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Input[str]:
        """
        FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter(name="controllerReport")
    def controller_report(self) -> Optional[pulumi.Input['Extender1ControllerReportArgs']]:
        """
        FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        """
        return pulumi.get(self, "controller_report")

    @controller_report.setter
    def controller_report(self, value: Optional[pulumi.Input['Extender1ControllerReportArgs']]):
        pulumi.set(self, "controller_report", value)

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
    @pulumi.getter(name="extName")
    def ext_name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender name.
        """
        return pulumi.get(self, "ext_name")

    @ext_name.setter
    def ext_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_name", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender serial number.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="loginPassword")
    def login_password(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender login password.
        """
        return pulumi.get(self, "login_password")

    @login_password.setter
    def login_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_password", value)

    @property
    @pulumi.getter
    def modem1(self) -> Optional[pulumi.Input['Extender1Modem1Args']]:
        """
        Configuration options for modem 1. The structure of `modem1` block is documented below.
        """
        return pulumi.get(self, "modem1")

    @modem1.setter
    def modem1(self, value: Optional[pulumi.Input['Extender1Modem1Args']]):
        pulumi.set(self, "modem1", value)

    @property
    @pulumi.getter
    def modem2(self) -> Optional[pulumi.Input['Extender1Modem2Args']]:
        """
        Configuration options for modem 2. The structure of `modem2` block is documented below.
        """
        return pulumi.get(self, "modem2")

    @modem2.setter
    def modem2(self, value: Optional[pulumi.Input['Extender1Modem2Args']]):
        pulumi.set(self, "modem2", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[int]]:
        """
        VDOM
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vdom", value)

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
class _Extender1State:
    def __init__(__self__, *,
                 authorized: Optional[pulumi.Input[str]] = None,
                 controller_report: Optional[pulumi.Input['Extender1ControllerReportArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_name: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 login_password: Optional[pulumi.Input[str]] = None,
                 modem1: Optional[pulumi.Input['Extender1Modem1Args']] = None,
                 modem2: Optional[pulumi.Input['Extender1Modem2Args']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Extender1 resources.
        :param pulumi.Input[str] authorized: FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        :param pulumi.Input['Extender1ControllerReportArgs'] controller_report: FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] ext_name: FortiExtender name.
        :param pulumi.Input[str] fosid: FortiExtender serial number.
        :param pulumi.Input[str] login_password: FortiExtender login password.
        :param pulumi.Input['Extender1Modem1Args'] modem1: Configuration options for modem 1. The structure of `modem1` block is documented below.
        :param pulumi.Input['Extender1Modem2Args'] modem2: Configuration options for modem 2. The structure of `modem2` block is documented below.
        :param pulumi.Input[str] name: FortiExtender entry name.
        :param pulumi.Input[int] vdom: VDOM
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authorized is not None:
            pulumi.set(__self__, "authorized", authorized)
        if controller_report is not None:
            pulumi.set(__self__, "controller_report", controller_report)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ext_name is not None:
            pulumi.set(__self__, "ext_name", ext_name)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if login_password is not None:
            pulumi.set(__self__, "login_password", login_password)
        if modem1 is not None:
            pulumi.set(__self__, "modem1", modem1)
        if modem2 is not None:
            pulumi.set(__self__, "modem2", modem2)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authorized(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter(name="controllerReport")
    def controller_report(self) -> Optional[pulumi.Input['Extender1ControllerReportArgs']]:
        """
        FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        """
        return pulumi.get(self, "controller_report")

    @controller_report.setter
    def controller_report(self, value: Optional[pulumi.Input['Extender1ControllerReportArgs']]):
        pulumi.set(self, "controller_report", value)

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
    @pulumi.getter(name="extName")
    def ext_name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender name.
        """
        return pulumi.get(self, "ext_name")

    @ext_name.setter
    def ext_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_name", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender serial number.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="loginPassword")
    def login_password(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender login password.
        """
        return pulumi.get(self, "login_password")

    @login_password.setter
    def login_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_password", value)

    @property
    @pulumi.getter
    def modem1(self) -> Optional[pulumi.Input['Extender1Modem1Args']]:
        """
        Configuration options for modem 1. The structure of `modem1` block is documented below.
        """
        return pulumi.get(self, "modem1")

    @modem1.setter
    def modem1(self, value: Optional[pulumi.Input['Extender1Modem1Args']]):
        pulumi.set(self, "modem1", value)

    @property
    @pulumi.getter
    def modem2(self) -> Optional[pulumi.Input['Extender1Modem2Args']]:
        """
        Configuration options for modem 2. The structure of `modem2` block is documented below.
        """
        return pulumi.get(self, "modem2")

    @modem2.setter
    def modem2(self, value: Optional[pulumi.Input['Extender1Modem2Args']]):
        pulumi.set(self, "modem2", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiExtender entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[int]]:
        """
        VDOM
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vdom", value)

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


class Extender1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[str]] = None,
                 controller_report: Optional[pulumi.Input[pulumi.InputType['Extender1ControllerReportArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_name: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 login_password: Optional[pulumi.Input[str]] = None,
                 modem1: Optional[pulumi.Input[pulumi.InputType['Extender1Modem1Args']]] = None,
                 modem2: Optional[pulumi.Input[pulumi.InputType['Extender1Modem2Args']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Extender controller configuration.

        > The resource applies to FortiOS Version >= 6.4.2. For FortiOS Version < 6.4.2, see `extendercontroller.Extender`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.extendercontroller.Extender1("trname",
            authorized="disable",
            controller_report=fortios.extendercontroller.Extender1ControllerReportArgs(
                interval=300,
                signal_threshold=10,
                status="disable",
            ),
            ext_name="2932",
            fosid="FX201E5919004031",
            modem1=fortios.extendercontroller.Extender1Modem1Args(
                auto_switch=fortios.extendercontroller.Extender1Modem1AutoSwitchArgs(
                    dataplan="disable",
                    disconnect="disable",
                    disconnect_period=600,
                    disconnect_threshold=3,
                    signal="disable",
                    switch_back="timer",
                    switch_back_time="00:01",
                    switch_back_timer=86400,
                ),
                conn_status=0,
                default_sim="sim2",
                gps="enable",
                redundant_intf="s1",
                redundant_mode="enable",
                sim1_pin="disable",
                sim1_pin_code="testpincode",
                sim2_pin="disable",
            ),
            modem2=fortios.extendercontroller.Extender1Modem2Args(
                auto_switch=fortios.extendercontroller.Extender1Modem2AutoSwitchArgs(
                    dataplan="disable",
                    disconnect="disable",
                    disconnect_period=600,
                    disconnect_threshold=3,
                    signal="disable",
                    switch_back_time="00:01",
                    switch_back_timer=86400,
                ),
                conn_status=0,
                default_sim="sim1",
                gps="enable",
                redundant_mode="disable",
                sim1_pin="disable",
                sim2_pin="disable",
            ),
            vdom=0)
        ```

        ## Import

        ExtenderController Extender1 can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized: FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        :param pulumi.Input[pulumi.InputType['Extender1ControllerReportArgs']] controller_report: FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] ext_name: FortiExtender name.
        :param pulumi.Input[str] fosid: FortiExtender serial number.
        :param pulumi.Input[str] login_password: FortiExtender login password.
        :param pulumi.Input[pulumi.InputType['Extender1Modem1Args']] modem1: Configuration options for modem 1. The structure of `modem1` block is documented below.
        :param pulumi.Input[pulumi.InputType['Extender1Modem2Args']] modem2: Configuration options for modem 2. The structure of `modem2` block is documented below.
        :param pulumi.Input[str] name: FortiExtender entry name.
        :param pulumi.Input[int] vdom: VDOM
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Extender1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Extender controller configuration.

        > The resource applies to FortiOS Version >= 6.4.2. For FortiOS Version < 6.4.2, see `extendercontroller.Extender`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.extendercontroller.Extender1("trname",
            authorized="disable",
            controller_report=fortios.extendercontroller.Extender1ControllerReportArgs(
                interval=300,
                signal_threshold=10,
                status="disable",
            ),
            ext_name="2932",
            fosid="FX201E5919004031",
            modem1=fortios.extendercontroller.Extender1Modem1Args(
                auto_switch=fortios.extendercontroller.Extender1Modem1AutoSwitchArgs(
                    dataplan="disable",
                    disconnect="disable",
                    disconnect_period=600,
                    disconnect_threshold=3,
                    signal="disable",
                    switch_back="timer",
                    switch_back_time="00:01",
                    switch_back_timer=86400,
                ),
                conn_status=0,
                default_sim="sim2",
                gps="enable",
                redundant_intf="s1",
                redundant_mode="enable",
                sim1_pin="disable",
                sim1_pin_code="testpincode",
                sim2_pin="disable",
            ),
            modem2=fortios.extendercontroller.Extender1Modem2Args(
                auto_switch=fortios.extendercontroller.Extender1Modem2AutoSwitchArgs(
                    dataplan="disable",
                    disconnect="disable",
                    disconnect_period=600,
                    disconnect_threshold=3,
                    signal="disable",
                    switch_back_time="00:01",
                    switch_back_timer=86400,
                ),
                conn_status=0,
                default_sim="sim1",
                gps="enable",
                redundant_mode="disable",
                sim1_pin="disable",
                sim2_pin="disable",
            ),
            vdom=0)
        ```

        ## Import

        ExtenderController Extender1 can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param Extender1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Extender1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[str]] = None,
                 controller_report: Optional[pulumi.Input[pulumi.InputType['Extender1ControllerReportArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_name: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 login_password: Optional[pulumi.Input[str]] = None,
                 modem1: Optional[pulumi.Input[pulumi.InputType['Extender1Modem1Args']]] = None,
                 modem2: Optional[pulumi.Input[pulumi.InputType['Extender1Modem2Args']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Extender1Args.__new__(Extender1Args)

            if authorized is None and not opts.urn:
                raise TypeError("Missing required property 'authorized'")
            __props__.__dict__["authorized"] = authorized
            __props__.__dict__["controller_report"] = controller_report
            __props__.__dict__["description"] = description
            __props__.__dict__["ext_name"] = ext_name
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["login_password"] = None if login_password is None else pulumi.Output.secret(login_password)
            __props__.__dict__["modem1"] = modem1
            __props__.__dict__["modem2"] = modem2
            __props__.__dict__["name"] = name
            __props__.__dict__["vdom"] = vdom
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["loginPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Extender1, __self__).__init__(
            'fortios:extendercontroller/extender1:Extender1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized: Optional[pulumi.Input[str]] = None,
            controller_report: Optional[pulumi.Input[pulumi.InputType['Extender1ControllerReportArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ext_name: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[str]] = None,
            login_password: Optional[pulumi.Input[str]] = None,
            modem1: Optional[pulumi.Input[pulumi.InputType['Extender1Modem1Args']]] = None,
            modem2: Optional[pulumi.Input[pulumi.InputType['Extender1Modem2Args']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdom: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Extender1':
        """
        Get an existing Extender1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized: FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        :param pulumi.Input[pulumi.InputType['Extender1ControllerReportArgs']] controller_report: FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] ext_name: FortiExtender name.
        :param pulumi.Input[str] fosid: FortiExtender serial number.
        :param pulumi.Input[str] login_password: FortiExtender login password.
        :param pulumi.Input[pulumi.InputType['Extender1Modem1Args']] modem1: Configuration options for modem 1. The structure of `modem1` block is documented below.
        :param pulumi.Input[pulumi.InputType['Extender1Modem2Args']] modem2: Configuration options for modem 2. The structure of `modem2` block is documented below.
        :param pulumi.Input[str] name: FortiExtender entry name.
        :param pulumi.Input[int] vdom: VDOM
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Extender1State.__new__(_Extender1State)

        __props__.__dict__["authorized"] = authorized
        __props__.__dict__["controller_report"] = controller_report
        __props__.__dict__["description"] = description
        __props__.__dict__["ext_name"] = ext_name
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["login_password"] = login_password
        __props__.__dict__["modem1"] = modem1
        __props__.__dict__["modem2"] = modem2
        __props__.__dict__["name"] = name
        __props__.__dict__["vdom"] = vdom
        __props__.__dict__["vdomparam"] = vdomparam
        return Extender1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Output[str]:
        """
        FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "authorized")

    @property
    @pulumi.getter(name="controllerReport")
    def controller_report(self) -> pulumi.Output['outputs.Extender1ControllerReport']:
        """
        FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        """
        return pulumi.get(self, "controller_report")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="extName")
    def ext_name(self) -> pulumi.Output[str]:
        """
        FortiExtender name.
        """
        return pulumi.get(self, "ext_name")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[str]:
        """
        FortiExtender serial number.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="loginPassword")
    def login_password(self) -> pulumi.Output[Optional[str]]:
        """
        FortiExtender login password.
        """
        return pulumi.get(self, "login_password")

    @property
    @pulumi.getter
    def modem1(self) -> pulumi.Output['outputs.Extender1Modem1']:
        """
        Configuration options for modem 1. The structure of `modem1` block is documented below.
        """
        return pulumi.get(self, "modem1")

    @property
    @pulumi.getter
    def modem2(self) -> pulumi.Output['outputs.Extender1Modem2']:
        """
        Configuration options for modem 2. The structure of `modem2` block is documented below.
        """
        return pulumi.get(self, "modem2")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FortiExtender entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[int]:
        """
        VDOM
        """
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

