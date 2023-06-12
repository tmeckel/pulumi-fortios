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

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input['ProfileFileFilterArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_credential_type: Optional[pulumi.Input[str]] = None,
                 server_keytabs: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[str] domain_controller: Domain for which to decrypt CIFS traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input['ProfileFileFilterArgs'] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] server_credential_type: CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]] server_keytabs: Server keytab. The structure of `server_keytab` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if domain_controller is not None:
            pulumi.set(__self__, "domain_controller", domain_controller)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if file_filter is not None:
            pulumi.set(__self__, "file_filter", file_filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_credential_type is not None:
            pulumi.set(__self__, "server_credential_type", server_credential_type)
        if server_keytabs is not None:
            pulumi.set(__self__, "server_keytabs", server_keytabs)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="domainController")
    def domain_controller(self) -> Optional[pulumi.Input[str]]:
        """
        Domain for which to decrypt CIFS traffic.
        """
        return pulumi.get(self, "domain_controller")

    @domain_controller.setter
    def domain_controller(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_controller", value)

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
    @pulumi.getter(name="fileFilter")
    def file_filter(self) -> Optional[pulumi.Input['ProfileFileFilterArgs']]:
        """
        File filter. The structure of `file_filter` block is documented below.
        """
        return pulumi.get(self, "file_filter")

    @file_filter.setter
    def file_filter(self, value: Optional[pulumi.Input['ProfileFileFilterArgs']]):
        pulumi.set(self, "file_filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverCredentialType")
    def server_credential_type(self) -> Optional[pulumi.Input[str]]:
        """
        CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        """
        return pulumi.get(self, "server_credential_type")

    @server_credential_type.setter
    def server_credential_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_credential_type", value)

    @property
    @pulumi.getter(name="serverKeytabs")
    def server_keytabs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]]]:
        """
        Server keytab. The structure of `server_keytab` block is documented below.
        """
        return pulumi.get(self, "server_keytabs")

    @server_keytabs.setter
    def server_keytabs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]]]):
        pulumi.set(self, "server_keytabs", value)

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
class _ProfileState:
    def __init__(__self__, *,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input['ProfileFileFilterArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_credential_type: Optional[pulumi.Input[str]] = None,
                 server_keytabs: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Profile resources.
        :param pulumi.Input[str] domain_controller: Domain for which to decrypt CIFS traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input['ProfileFileFilterArgs'] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] server_credential_type: CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]] server_keytabs: Server keytab. The structure of `server_keytab` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if domain_controller is not None:
            pulumi.set(__self__, "domain_controller", domain_controller)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if file_filter is not None:
            pulumi.set(__self__, "file_filter", file_filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_credential_type is not None:
            pulumi.set(__self__, "server_credential_type", server_credential_type)
        if server_keytabs is not None:
            pulumi.set(__self__, "server_keytabs", server_keytabs)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="domainController")
    def domain_controller(self) -> Optional[pulumi.Input[str]]:
        """
        Domain for which to decrypt CIFS traffic.
        """
        return pulumi.get(self, "domain_controller")

    @domain_controller.setter
    def domain_controller(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_controller", value)

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
    @pulumi.getter(name="fileFilter")
    def file_filter(self) -> Optional[pulumi.Input['ProfileFileFilterArgs']]:
        """
        File filter. The structure of `file_filter` block is documented below.
        """
        return pulumi.get(self, "file_filter")

    @file_filter.setter
    def file_filter(self, value: Optional[pulumi.Input['ProfileFileFilterArgs']]):
        pulumi.set(self, "file_filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverCredentialType")
    def server_credential_type(self) -> Optional[pulumi.Input[str]]:
        """
        CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        """
        return pulumi.get(self, "server_credential_type")

    @server_credential_type.setter
    def server_credential_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_credential_type", value)

    @property
    @pulumi.getter(name="serverKeytabs")
    def server_keytabs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]]]:
        """
        Server keytab. The structure of `server_keytab` block is documented below.
        """
        return pulumi.get(self, "server_keytabs")

    @server_keytabs.setter
    def server_keytabs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileServerKeytabArgs']]]]):
        pulumi.set(self, "server_keytabs", value)

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


class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input[pulumi.InputType['ProfileFileFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_credential_type: Optional[pulumi.Input[str]] = None,
                 server_keytabs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileServerKeytabArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure CIFS profile. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1`.

        ## Import

        Cifs Profile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:cifs/profile:Profile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:cifs/profile:Profile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_controller: Domain for which to decrypt CIFS traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[pulumi.InputType['ProfileFileFilterArgs']] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] server_credential_type: CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileServerKeytabArgs']]]] server_keytabs: Server keytab. The structure of `server_keytab` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure CIFS profile. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1`.

        ## Import

        Cifs Profile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:cifs/profile:Profile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:cifs/profile:Profile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input[pulumi.InputType['ProfileFileFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_credential_type: Optional[pulumi.Input[str]] = None,
                 server_keytabs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileServerKeytabArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["domain_controller"] = domain_controller
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["file_filter"] = file_filter
            __props__.__dict__["name"] = name
            __props__.__dict__["server_credential_type"] = server_credential_type
            __props__.__dict__["server_keytabs"] = server_keytabs
            __props__.__dict__["vdomparam"] = vdomparam
        super(Profile, __self__).__init__(
            'fortios:cifs/profile:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_controller: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            file_filter: Optional[pulumi.Input[pulumi.InputType['ProfileFileFilterArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            server_credential_type: Optional[pulumi.Input[str]] = None,
            server_keytabs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileServerKeytabArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_controller: Domain for which to decrypt CIFS traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[pulumi.InputType['ProfileFileFilterArgs']] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] server_credential_type: CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileServerKeytabArgs']]]] server_keytabs: Server keytab. The structure of `server_keytab` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProfileState.__new__(_ProfileState)

        __props__.__dict__["domain_controller"] = domain_controller
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["file_filter"] = file_filter
        __props__.__dict__["name"] = name
        __props__.__dict__["server_credential_type"] = server_credential_type
        __props__.__dict__["server_keytabs"] = server_keytabs
        __props__.__dict__["vdomparam"] = vdomparam
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainController")
    def domain_controller(self) -> pulumi.Output[str]:
        """
        Domain for which to decrypt CIFS traffic.
        """
        return pulumi.get(self, "domain_controller")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="fileFilter")
    def file_filter(self) -> pulumi.Output['outputs.ProfileFileFilter']:
        """
        File filter. The structure of `file_filter` block is documented below.
        """
        return pulumi.get(self, "file_filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverCredentialType")
    def server_credential_type(self) -> pulumi.Output[str]:
        """
        CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
        """
        return pulumi.get(self, "server_credential_type")

    @property
    @pulumi.getter(name="serverKeytabs")
    def server_keytabs(self) -> pulumi.Output[Optional[Sequence['outputs.ProfileServerKeytab']]]:
        """
        Server keytab. The structure of `server_keytab` block is documented below.
        """
        return pulumi.get(self, "server_keytabs")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
