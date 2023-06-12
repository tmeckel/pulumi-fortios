# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CpusArgs', 'Cpus']

@pulumi.input_type
class CpusArgs:
    def __init__(__self__, *,
                 ips_cpus: Optional[pulumi.Input[str]] = None,
                 isolated_cpus: Optional[pulumi.Input[str]] = None,
                 rx_cpus: Optional[pulumi.Input[str]] = None,
                 tx_cpus: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vnp_cpus: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cpus resource.
        :param pulumi.Input[str] ips_cpus: CPUs enabled to run DPDK IPS engines.
        :param pulumi.Input[str] isolated_cpus: CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        :param pulumi.Input[str] rx_cpus: CPUs enabled to run DPDK RX engines.
        :param pulumi.Input[str] tx_cpus: CPUs enabled to run DPDK TX engines.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vnp_cpus: CPUs enabled to run DPDK VNP engines.
        """
        if ips_cpus is not None:
            pulumi.set(__self__, "ips_cpus", ips_cpus)
        if isolated_cpus is not None:
            pulumi.set(__self__, "isolated_cpus", isolated_cpus)
        if rx_cpus is not None:
            pulumi.set(__self__, "rx_cpus", rx_cpus)
        if tx_cpus is not None:
            pulumi.set(__self__, "tx_cpus", tx_cpus)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vnp_cpus is not None:
            pulumi.set(__self__, "vnp_cpus", vnp_cpus)

    @property
    @pulumi.getter(name="ipsCpus")
    def ips_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK IPS engines.
        """
        return pulumi.get(self, "ips_cpus")

    @ips_cpus.setter
    def ips_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ips_cpus", value)

    @property
    @pulumi.getter(name="isolatedCpus")
    def isolated_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        """
        return pulumi.get(self, "isolated_cpus")

    @isolated_cpus.setter
    def isolated_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isolated_cpus", value)

    @property
    @pulumi.getter(name="rxCpus")
    def rx_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK RX engines.
        """
        return pulumi.get(self, "rx_cpus")

    @rx_cpus.setter
    def rx_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rx_cpus", value)

    @property
    @pulumi.getter(name="txCpus")
    def tx_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK TX engines.
        """
        return pulumi.get(self, "tx_cpus")

    @tx_cpus.setter
    def tx_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tx_cpus", value)

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
    @pulumi.getter(name="vnpCpus")
    def vnp_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK VNP engines.
        """
        return pulumi.get(self, "vnp_cpus")

    @vnp_cpus.setter
    def vnp_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vnp_cpus", value)


@pulumi.input_type
class _CpusState:
    def __init__(__self__, *,
                 ips_cpus: Optional[pulumi.Input[str]] = None,
                 isolated_cpus: Optional[pulumi.Input[str]] = None,
                 rx_cpus: Optional[pulumi.Input[str]] = None,
                 tx_cpus: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vnp_cpus: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cpus resources.
        :param pulumi.Input[str] ips_cpus: CPUs enabled to run DPDK IPS engines.
        :param pulumi.Input[str] isolated_cpus: CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        :param pulumi.Input[str] rx_cpus: CPUs enabled to run DPDK RX engines.
        :param pulumi.Input[str] tx_cpus: CPUs enabled to run DPDK TX engines.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vnp_cpus: CPUs enabled to run DPDK VNP engines.
        """
        if ips_cpus is not None:
            pulumi.set(__self__, "ips_cpus", ips_cpus)
        if isolated_cpus is not None:
            pulumi.set(__self__, "isolated_cpus", isolated_cpus)
        if rx_cpus is not None:
            pulumi.set(__self__, "rx_cpus", rx_cpus)
        if tx_cpus is not None:
            pulumi.set(__self__, "tx_cpus", tx_cpus)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vnp_cpus is not None:
            pulumi.set(__self__, "vnp_cpus", vnp_cpus)

    @property
    @pulumi.getter(name="ipsCpus")
    def ips_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK IPS engines.
        """
        return pulumi.get(self, "ips_cpus")

    @ips_cpus.setter
    def ips_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ips_cpus", value)

    @property
    @pulumi.getter(name="isolatedCpus")
    def isolated_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        """
        return pulumi.get(self, "isolated_cpus")

    @isolated_cpus.setter
    def isolated_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isolated_cpus", value)

    @property
    @pulumi.getter(name="rxCpus")
    def rx_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK RX engines.
        """
        return pulumi.get(self, "rx_cpus")

    @rx_cpus.setter
    def rx_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rx_cpus", value)

    @property
    @pulumi.getter(name="txCpus")
    def tx_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK TX engines.
        """
        return pulumi.get(self, "tx_cpus")

    @tx_cpus.setter
    def tx_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tx_cpus", value)

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
    @pulumi.getter(name="vnpCpus")
    def vnp_cpus(self) -> Optional[pulumi.Input[str]]:
        """
        CPUs enabled to run DPDK VNP engines.
        """
        return pulumi.get(self, "vnp_cpus")

    @vnp_cpus.setter
    def vnp_cpus(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vnp_cpus", value)


class Cpus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ips_cpus: Optional[pulumi.Input[str]] = None,
                 isolated_cpus: Optional[pulumi.Input[str]] = None,
                 rx_cpus: Optional[pulumi.Input[str]] = None,
                 tx_cpus: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vnp_cpus: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure CPUs enabled to run engines in each DPDK stage. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Dpdk Cpus can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:dpdk/cpus:Cpus labelname DpdkCpus
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:dpdk/cpus:Cpus labelname DpdkCpus
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ips_cpus: CPUs enabled to run DPDK IPS engines.
        :param pulumi.Input[str] isolated_cpus: CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        :param pulumi.Input[str] rx_cpus: CPUs enabled to run DPDK RX engines.
        :param pulumi.Input[str] tx_cpus: CPUs enabled to run DPDK TX engines.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vnp_cpus: CPUs enabled to run DPDK VNP engines.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CpusArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure CPUs enabled to run engines in each DPDK stage. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        Dpdk Cpus can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:dpdk/cpus:Cpus labelname DpdkCpus
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:dpdk/cpus:Cpus labelname DpdkCpus
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param CpusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CpusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ips_cpus: Optional[pulumi.Input[str]] = None,
                 isolated_cpus: Optional[pulumi.Input[str]] = None,
                 rx_cpus: Optional[pulumi.Input[str]] = None,
                 tx_cpus: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vnp_cpus: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CpusArgs.__new__(CpusArgs)

            __props__.__dict__["ips_cpus"] = ips_cpus
            __props__.__dict__["isolated_cpus"] = isolated_cpus
            __props__.__dict__["rx_cpus"] = rx_cpus
            __props__.__dict__["tx_cpus"] = tx_cpus
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vnp_cpus"] = vnp_cpus
        super(Cpus, __self__).__init__(
            'fortios:dpdk/cpus:Cpus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ips_cpus: Optional[pulumi.Input[str]] = None,
            isolated_cpus: Optional[pulumi.Input[str]] = None,
            rx_cpus: Optional[pulumi.Input[str]] = None,
            tx_cpus: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vnp_cpus: Optional[pulumi.Input[str]] = None) -> 'Cpus':
        """
        Get an existing Cpus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ips_cpus: CPUs enabled to run DPDK IPS engines.
        :param pulumi.Input[str] isolated_cpus: CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        :param pulumi.Input[str] rx_cpus: CPUs enabled to run DPDK RX engines.
        :param pulumi.Input[str] tx_cpus: CPUs enabled to run DPDK TX engines.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vnp_cpus: CPUs enabled to run DPDK VNP engines.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CpusState.__new__(_CpusState)

        __props__.__dict__["ips_cpus"] = ips_cpus
        __props__.__dict__["isolated_cpus"] = isolated_cpus
        __props__.__dict__["rx_cpus"] = rx_cpus
        __props__.__dict__["tx_cpus"] = tx_cpus
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vnp_cpus"] = vnp_cpus
        return Cpus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipsCpus")
    def ips_cpus(self) -> pulumi.Output[str]:
        """
        CPUs enabled to run DPDK IPS engines.
        """
        return pulumi.get(self, "ips_cpus")

    @property
    @pulumi.getter(name="isolatedCpus")
    def isolated_cpus(self) -> pulumi.Output[str]:
        """
        CPUs isolated to run only the DPDK engines with the exception of processes that have affinity explicitly set by either a user configuration or by their implementation.
        """
        return pulumi.get(self, "isolated_cpus")

    @property
    @pulumi.getter(name="rxCpus")
    def rx_cpus(self) -> pulumi.Output[str]:
        """
        CPUs enabled to run DPDK RX engines.
        """
        return pulumi.get(self, "rx_cpus")

    @property
    @pulumi.getter(name="txCpus")
    def tx_cpus(self) -> pulumi.Output[str]:
        """
        CPUs enabled to run DPDK TX engines.
        """
        return pulumi.get(self, "tx_cpus")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="vnpCpus")
    def vnp_cpus(self) -> pulumi.Output[str]:
        """
        CPUs enabled to run DPDK VNP engines.
        """
        return pulumi.get(self, "vnp_cpus")

