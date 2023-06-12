# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CentralsnatmapMoveArgs', 'CentralsnatmapMove']

@pulumi.input_type
class CentralsnatmapMoveArgs:
    def __init__(__self__, *,
                 move: pulumi.Input[str],
                 policyid_dst: pulumi.Input[int],
                 policyid_src: pulumi.Input[int],
                 comment: Optional[pulumi.Input[str]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CentralsnatmapMove resource.
        """
        pulumi.set(__self__, "move", move)
        pulumi.set(__self__, "policyid_dst", policyid_dst)
        pulumi.set(__self__, "policyid_src", policyid_src)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if state_policy_srcdst_pos is not None:
            pulumi.set(__self__, "state_policy_srcdst_pos", state_policy_srcdst_pos)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def move(self) -> pulumi.Input[str]:
        return pulumi.get(self, "move")

    @move.setter
    def move(self, value: pulumi.Input[str]):
        pulumi.set(self, "move", value)

    @property
    @pulumi.getter(name="policyidDst")
    def policyid_dst(self) -> pulumi.Input[int]:
        return pulumi.get(self, "policyid_dst")

    @policyid_dst.setter
    def policyid_dst(self, value: pulumi.Input[int]):
        pulumi.set(self, "policyid_dst", value)

    @property
    @pulumi.getter(name="policyidSrc")
    def policyid_src(self) -> pulumi.Input[int]:
        return pulumi.get(self, "policyid_src")

    @policyid_src.setter
    def policyid_src(self, value: pulumi.Input[int]):
        pulumi.set(self, "policyid_src", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="statePolicySrcdstPos")
    def state_policy_srcdst_pos(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state_policy_srcdst_pos")

    @state_policy_srcdst_pos.setter
    def state_policy_srcdst_pos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_policy_srcdst_pos", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _CentralsnatmapMoveState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 move: Optional[pulumi.Input[str]] = None,
                 policyid_dst: Optional[pulumi.Input[int]] = None,
                 policyid_src: Optional[pulumi.Input[int]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CentralsnatmapMove resources.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if move is not None:
            pulumi.set(__self__, "move", move)
        if policyid_dst is not None:
            pulumi.set(__self__, "policyid_dst", policyid_dst)
        if policyid_src is not None:
            pulumi.set(__self__, "policyid_src", policyid_src)
        if state_policy_srcdst_pos is not None:
            pulumi.set(__self__, "state_policy_srcdst_pos", state_policy_srcdst_pos)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def move(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "move")

    @move.setter
    def move(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "move", value)

    @property
    @pulumi.getter(name="policyidDst")
    def policyid_dst(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policyid_dst")

    @policyid_dst.setter
    def policyid_dst(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policyid_dst", value)

    @property
    @pulumi.getter(name="policyidSrc")
    def policyid_src(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policyid_src")

    @policyid_src.setter
    def policyid_src(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policyid_src", value)

    @property
    @pulumi.getter(name="statePolicySrcdstPos")
    def state_policy_srcdst_pos(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state_policy_srcdst_pos")

    @state_policy_srcdst_pos.setter
    def state_policy_srcdst_pos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_policy_srcdst_pos", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class CentralsnatmapMove(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 move: Optional[pulumi.Input[str]] = None,
                 policyid_dst: Optional[pulumi.Input[int]] = None,
                 policyid_src: Optional[pulumi.Input[int]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CentralsnatmapMove resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CentralsnatmapMoveArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CentralsnatmapMove resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CentralsnatmapMoveArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CentralsnatmapMoveArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 move: Optional[pulumi.Input[str]] = None,
                 policyid_dst: Optional[pulumi.Input[int]] = None,
                 policyid_src: Optional[pulumi.Input[int]] = None,
                 state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CentralsnatmapMoveArgs.__new__(CentralsnatmapMoveArgs)

            __props__.__dict__["comment"] = comment
            if move is None and not opts.urn:
                raise TypeError("Missing required property 'move'")
            __props__.__dict__["move"] = move
            if policyid_dst is None and not opts.urn:
                raise TypeError("Missing required property 'policyid_dst'")
            __props__.__dict__["policyid_dst"] = policyid_dst
            if policyid_src is None and not opts.urn:
                raise TypeError("Missing required property 'policyid_src'")
            __props__.__dict__["policyid_src"] = policyid_src
            __props__.__dict__["state_policy_srcdst_pos"] = state_policy_srcdst_pos
            __props__.__dict__["vdomparam"] = vdomparam
        super(CentralsnatmapMove, __self__).__init__(
            'fortios:firewall/centralsnatmapMove:CentralsnatmapMove',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            move: Optional[pulumi.Input[str]] = None,
            policyid_dst: Optional[pulumi.Input[int]] = None,
            policyid_src: Optional[pulumi.Input[int]] = None,
            state_policy_srcdst_pos: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'CentralsnatmapMove':
        """
        Get an existing CentralsnatmapMove resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CentralsnatmapMoveState.__new__(_CentralsnatmapMoveState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["move"] = move
        __props__.__dict__["policyid_dst"] = policyid_dst
        __props__.__dict__["policyid_src"] = policyid_src
        __props__.__dict__["state_policy_srcdst_pos"] = state_policy_srcdst_pos
        __props__.__dict__["vdomparam"] = vdomparam
        return CentralsnatmapMove(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def move(self) -> pulumi.Output[str]:
        return pulumi.get(self, "move")

    @property
    @pulumi.getter(name="policyidDst")
    def policyid_dst(self) -> pulumi.Output[int]:
        return pulumi.get(self, "policyid_dst")

    @property
    @pulumi.getter(name="policyidSrc")
    def policyid_src(self) -> pulumi.Output[int]:
        return pulumi.get(self, "policyid_src")

    @property
    @pulumi.getter(name="statePolicySrcdstPos")
    def state_policy_srcdst_pos(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "state_policy_srcdst_pos")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

