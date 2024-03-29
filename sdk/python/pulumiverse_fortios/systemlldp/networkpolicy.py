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

__all__ = ['NetworkpolicyArgs', 'Networkpolicy']

@pulumi.input_type
class NetworkpolicyArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input['NetworkpolicyGuestArgs']] = None,
                 guest_voice_signaling: Optional[pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input['NetworkpolicySoftphoneArgs']] = None,
                 streaming_video: Optional[pulumi.Input['NetworkpolicyStreamingVideoArgs']] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input['NetworkpolicyVideoConferencingArgs']] = None,
                 video_signaling: Optional[pulumi.Input['NetworkpolicyVideoSignalingArgs']] = None,
                 voice: Optional[pulumi.Input['NetworkpolicyVoiceArgs']] = None,
                 voice_signaling: Optional[pulumi.Input['NetworkpolicyVoiceSignalingArgs']] = None):
        """
        The set of arguments for constructing a Networkpolicy resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input['NetworkpolicyGuestArgs'] guest: Guest. The structure of `guest` block is documented below.
        :param pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs'] guest_voice_signaling: Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        :param pulumi.Input[str] name: LLDP network policy name.
        :param pulumi.Input['NetworkpolicySoftphoneArgs'] softphone: Softphone. The structure of `softphone` block is documented below.
        :param pulumi.Input['NetworkpolicyStreamingVideoArgs'] streaming_video: Streaming Video. The structure of `streaming_video` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input['NetworkpolicyVideoConferencingArgs'] video_conferencing: Video Conferencing. The structure of `video_conferencing` block is documented below.
        :param pulumi.Input['NetworkpolicyVideoSignalingArgs'] video_signaling: Video Signaling. The structure of `video_signaling` block is documented below.
        :param pulumi.Input['NetworkpolicyVoiceArgs'] voice: Voice. The structure of `voice` block is documented below.
        :param pulumi.Input['NetworkpolicyVoiceSignalingArgs'] voice_signaling: Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if guest is not None:
            pulumi.set(__self__, "guest", guest)
        if guest_voice_signaling is not None:
            pulumi.set(__self__, "guest_voice_signaling", guest_voice_signaling)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if softphone is not None:
            pulumi.set(__self__, "softphone", softphone)
        if streaming_video is not None:
            pulumi.set(__self__, "streaming_video", streaming_video)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if video_conferencing is not None:
            pulumi.set(__self__, "video_conferencing", video_conferencing)
        if video_signaling is not None:
            pulumi.set(__self__, "video_signaling", video_signaling)
        if voice is not None:
            pulumi.set(__self__, "voice", voice)
        if voice_signaling is not None:
            pulumi.set(__self__, "voice_signaling", voice_signaling)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def guest(self) -> Optional[pulumi.Input['NetworkpolicyGuestArgs']]:
        """
        Guest. The structure of `guest` block is documented below.
        """
        return pulumi.get(self, "guest")

    @guest.setter
    def guest(self, value: Optional[pulumi.Input['NetworkpolicyGuestArgs']]):
        pulumi.set(self, "guest", value)

    @property
    @pulumi.getter(name="guestVoiceSignaling")
    def guest_voice_signaling(self) -> Optional[pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs']]:
        """
        Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        """
        return pulumi.get(self, "guest_voice_signaling")

    @guest_voice_signaling.setter
    def guest_voice_signaling(self, value: Optional[pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs']]):
        pulumi.set(self, "guest_voice_signaling", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        LLDP network policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def softphone(self) -> Optional[pulumi.Input['NetworkpolicySoftphoneArgs']]:
        """
        Softphone. The structure of `softphone` block is documented below.
        """
        return pulumi.get(self, "softphone")

    @softphone.setter
    def softphone(self, value: Optional[pulumi.Input['NetworkpolicySoftphoneArgs']]):
        pulumi.set(self, "softphone", value)

    @property
    @pulumi.getter(name="streamingVideo")
    def streaming_video(self) -> Optional[pulumi.Input['NetworkpolicyStreamingVideoArgs']]:
        """
        Streaming Video. The structure of `streaming_video` block is documented below.
        """
        return pulumi.get(self, "streaming_video")

    @streaming_video.setter
    def streaming_video(self, value: Optional[pulumi.Input['NetworkpolicyStreamingVideoArgs']]):
        pulumi.set(self, "streaming_video", value)

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
    @pulumi.getter(name="videoConferencing")
    def video_conferencing(self) -> Optional[pulumi.Input['NetworkpolicyVideoConferencingArgs']]:
        """
        Video Conferencing. The structure of `video_conferencing` block is documented below.
        """
        return pulumi.get(self, "video_conferencing")

    @video_conferencing.setter
    def video_conferencing(self, value: Optional[pulumi.Input['NetworkpolicyVideoConferencingArgs']]):
        pulumi.set(self, "video_conferencing", value)

    @property
    @pulumi.getter(name="videoSignaling")
    def video_signaling(self) -> Optional[pulumi.Input['NetworkpolicyVideoSignalingArgs']]:
        """
        Video Signaling. The structure of `video_signaling` block is documented below.
        """
        return pulumi.get(self, "video_signaling")

    @video_signaling.setter
    def video_signaling(self, value: Optional[pulumi.Input['NetworkpolicyVideoSignalingArgs']]):
        pulumi.set(self, "video_signaling", value)

    @property
    @pulumi.getter
    def voice(self) -> Optional[pulumi.Input['NetworkpolicyVoiceArgs']]:
        """
        Voice. The structure of `voice` block is documented below.
        """
        return pulumi.get(self, "voice")

    @voice.setter
    def voice(self, value: Optional[pulumi.Input['NetworkpolicyVoiceArgs']]):
        pulumi.set(self, "voice", value)

    @property
    @pulumi.getter(name="voiceSignaling")
    def voice_signaling(self) -> Optional[pulumi.Input['NetworkpolicyVoiceSignalingArgs']]:
        """
        Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        return pulumi.get(self, "voice_signaling")

    @voice_signaling.setter
    def voice_signaling(self, value: Optional[pulumi.Input['NetworkpolicyVoiceSignalingArgs']]):
        pulumi.set(self, "voice_signaling", value)


@pulumi.input_type
class _NetworkpolicyState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input['NetworkpolicyGuestArgs']] = None,
                 guest_voice_signaling: Optional[pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input['NetworkpolicySoftphoneArgs']] = None,
                 streaming_video: Optional[pulumi.Input['NetworkpolicyStreamingVideoArgs']] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input['NetworkpolicyVideoConferencingArgs']] = None,
                 video_signaling: Optional[pulumi.Input['NetworkpolicyVideoSignalingArgs']] = None,
                 voice: Optional[pulumi.Input['NetworkpolicyVoiceArgs']] = None,
                 voice_signaling: Optional[pulumi.Input['NetworkpolicyVoiceSignalingArgs']] = None):
        """
        Input properties used for looking up and filtering Networkpolicy resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input['NetworkpolicyGuestArgs'] guest: Guest. The structure of `guest` block is documented below.
        :param pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs'] guest_voice_signaling: Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        :param pulumi.Input[str] name: LLDP network policy name.
        :param pulumi.Input['NetworkpolicySoftphoneArgs'] softphone: Softphone. The structure of `softphone` block is documented below.
        :param pulumi.Input['NetworkpolicyStreamingVideoArgs'] streaming_video: Streaming Video. The structure of `streaming_video` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input['NetworkpolicyVideoConferencingArgs'] video_conferencing: Video Conferencing. The structure of `video_conferencing` block is documented below.
        :param pulumi.Input['NetworkpolicyVideoSignalingArgs'] video_signaling: Video Signaling. The structure of `video_signaling` block is documented below.
        :param pulumi.Input['NetworkpolicyVoiceArgs'] voice: Voice. The structure of `voice` block is documented below.
        :param pulumi.Input['NetworkpolicyVoiceSignalingArgs'] voice_signaling: Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if guest is not None:
            pulumi.set(__self__, "guest", guest)
        if guest_voice_signaling is not None:
            pulumi.set(__self__, "guest_voice_signaling", guest_voice_signaling)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if softphone is not None:
            pulumi.set(__self__, "softphone", softphone)
        if streaming_video is not None:
            pulumi.set(__self__, "streaming_video", streaming_video)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if video_conferencing is not None:
            pulumi.set(__self__, "video_conferencing", video_conferencing)
        if video_signaling is not None:
            pulumi.set(__self__, "video_signaling", video_signaling)
        if voice is not None:
            pulumi.set(__self__, "voice", voice)
        if voice_signaling is not None:
            pulumi.set(__self__, "voice_signaling", voice_signaling)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def guest(self) -> Optional[pulumi.Input['NetworkpolicyGuestArgs']]:
        """
        Guest. The structure of `guest` block is documented below.
        """
        return pulumi.get(self, "guest")

    @guest.setter
    def guest(self, value: Optional[pulumi.Input['NetworkpolicyGuestArgs']]):
        pulumi.set(self, "guest", value)

    @property
    @pulumi.getter(name="guestVoiceSignaling")
    def guest_voice_signaling(self) -> Optional[pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs']]:
        """
        Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        """
        return pulumi.get(self, "guest_voice_signaling")

    @guest_voice_signaling.setter
    def guest_voice_signaling(self, value: Optional[pulumi.Input['NetworkpolicyGuestVoiceSignalingArgs']]):
        pulumi.set(self, "guest_voice_signaling", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        LLDP network policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def softphone(self) -> Optional[pulumi.Input['NetworkpolicySoftphoneArgs']]:
        """
        Softphone. The structure of `softphone` block is documented below.
        """
        return pulumi.get(self, "softphone")

    @softphone.setter
    def softphone(self, value: Optional[pulumi.Input['NetworkpolicySoftphoneArgs']]):
        pulumi.set(self, "softphone", value)

    @property
    @pulumi.getter(name="streamingVideo")
    def streaming_video(self) -> Optional[pulumi.Input['NetworkpolicyStreamingVideoArgs']]:
        """
        Streaming Video. The structure of `streaming_video` block is documented below.
        """
        return pulumi.get(self, "streaming_video")

    @streaming_video.setter
    def streaming_video(self, value: Optional[pulumi.Input['NetworkpolicyStreamingVideoArgs']]):
        pulumi.set(self, "streaming_video", value)

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
    @pulumi.getter(name="videoConferencing")
    def video_conferencing(self) -> Optional[pulumi.Input['NetworkpolicyVideoConferencingArgs']]:
        """
        Video Conferencing. The structure of `video_conferencing` block is documented below.
        """
        return pulumi.get(self, "video_conferencing")

    @video_conferencing.setter
    def video_conferencing(self, value: Optional[pulumi.Input['NetworkpolicyVideoConferencingArgs']]):
        pulumi.set(self, "video_conferencing", value)

    @property
    @pulumi.getter(name="videoSignaling")
    def video_signaling(self) -> Optional[pulumi.Input['NetworkpolicyVideoSignalingArgs']]:
        """
        Video Signaling. The structure of `video_signaling` block is documented below.
        """
        return pulumi.get(self, "video_signaling")

    @video_signaling.setter
    def video_signaling(self, value: Optional[pulumi.Input['NetworkpolicyVideoSignalingArgs']]):
        pulumi.set(self, "video_signaling", value)

    @property
    @pulumi.getter
    def voice(self) -> Optional[pulumi.Input['NetworkpolicyVoiceArgs']]:
        """
        Voice. The structure of `voice` block is documented below.
        """
        return pulumi.get(self, "voice")

    @voice.setter
    def voice(self, value: Optional[pulumi.Input['NetworkpolicyVoiceArgs']]):
        pulumi.set(self, "voice", value)

    @property
    @pulumi.getter(name="voiceSignaling")
    def voice_signaling(self) -> Optional[pulumi.Input['NetworkpolicyVoiceSignalingArgs']]:
        """
        Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        return pulumi.get(self, "voice_signaling")

    @voice_signaling.setter
    def voice_signaling(self, value: Optional[pulumi.Input['NetworkpolicyVoiceSignalingArgs']]):
        pulumi.set(self, "voice_signaling", value)


class Networkpolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyGuestArgs']]] = None,
                 guest_voice_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyGuestVoiceSignalingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input[pulumi.InputType['NetworkpolicySoftphoneArgs']]] = None,
                 streaming_video: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyStreamingVideoArgs']]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVideoConferencingArgs']]] = None,
                 video_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVideoSignalingArgs']]] = None,
                 voice: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVoiceArgs']]] = None,
                 voice_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVoiceSignalingArgs']]] = None,
                 __props__=None):
        """
        Configure LLDP network policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.systemlldp.Networkpolicy("trname", comment="test")
        ```

        ## Import

        SystemLldp NetworkPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:systemlldp/networkpolicy:Networkpolicy labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:systemlldp/networkpolicy:Networkpolicy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyGuestArgs']] guest: Guest. The structure of `guest` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyGuestVoiceSignalingArgs']] guest_voice_signaling: Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        :param pulumi.Input[str] name: LLDP network policy name.
        :param pulumi.Input[pulumi.InputType['NetworkpolicySoftphoneArgs']] softphone: Softphone. The structure of `softphone` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyStreamingVideoArgs']] streaming_video: Streaming Video. The structure of `streaming_video` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVideoConferencingArgs']] video_conferencing: Video Conferencing. The structure of `video_conferencing` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVideoSignalingArgs']] video_signaling: Video Signaling. The structure of `video_signaling` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVoiceArgs']] voice: Voice. The structure of `voice` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVoiceSignalingArgs']] voice_signaling: Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkpolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure LLDP network policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.systemlldp.Networkpolicy("trname", comment="test")
        ```

        ## Import

        SystemLldp NetworkPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:systemlldp/networkpolicy:Networkpolicy labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:systemlldp/networkpolicy:Networkpolicy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param NetworkpolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkpolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 guest: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyGuestArgs']]] = None,
                 guest_voice_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyGuestVoiceSignalingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 softphone: Optional[pulumi.Input[pulumi.InputType['NetworkpolicySoftphoneArgs']]] = None,
                 streaming_video: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyStreamingVideoArgs']]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 video_conferencing: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVideoConferencingArgs']]] = None,
                 video_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVideoSignalingArgs']]] = None,
                 voice: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVoiceArgs']]] = None,
                 voice_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVoiceSignalingArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkpolicyArgs.__new__(NetworkpolicyArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["guest"] = guest
            __props__.__dict__["guest_voice_signaling"] = guest_voice_signaling
            __props__.__dict__["name"] = name
            __props__.__dict__["softphone"] = softphone
            __props__.__dict__["streaming_video"] = streaming_video
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["video_conferencing"] = video_conferencing
            __props__.__dict__["video_signaling"] = video_signaling
            __props__.__dict__["voice"] = voice
            __props__.__dict__["voice_signaling"] = voice_signaling
        super(Networkpolicy, __self__).__init__(
            'fortios:systemlldp/networkpolicy:Networkpolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            guest: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyGuestArgs']]] = None,
            guest_voice_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyGuestVoiceSignalingArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            softphone: Optional[pulumi.Input[pulumi.InputType['NetworkpolicySoftphoneArgs']]] = None,
            streaming_video: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyStreamingVideoArgs']]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            video_conferencing: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVideoConferencingArgs']]] = None,
            video_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVideoSignalingArgs']]] = None,
            voice: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVoiceArgs']]] = None,
            voice_signaling: Optional[pulumi.Input[pulumi.InputType['NetworkpolicyVoiceSignalingArgs']]] = None) -> 'Networkpolicy':
        """
        Get an existing Networkpolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyGuestArgs']] guest: Guest. The structure of `guest` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyGuestVoiceSignalingArgs']] guest_voice_signaling: Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        :param pulumi.Input[str] name: LLDP network policy name.
        :param pulumi.Input[pulumi.InputType['NetworkpolicySoftphoneArgs']] softphone: Softphone. The structure of `softphone` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyStreamingVideoArgs']] streaming_video: Streaming Video. The structure of `streaming_video` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVideoConferencingArgs']] video_conferencing: Video Conferencing. The structure of `video_conferencing` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVideoSignalingArgs']] video_signaling: Video Signaling. The structure of `video_signaling` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVoiceArgs']] voice: Voice. The structure of `voice` block is documented below.
        :param pulumi.Input[pulumi.InputType['NetworkpolicyVoiceSignalingArgs']] voice_signaling: Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkpolicyState.__new__(_NetworkpolicyState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["guest"] = guest
        __props__.__dict__["guest_voice_signaling"] = guest_voice_signaling
        __props__.__dict__["name"] = name
        __props__.__dict__["softphone"] = softphone
        __props__.__dict__["streaming_video"] = streaming_video
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["video_conferencing"] = video_conferencing
        __props__.__dict__["video_signaling"] = video_signaling
        __props__.__dict__["voice"] = voice
        __props__.__dict__["voice_signaling"] = voice_signaling
        return Networkpolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def guest(self) -> pulumi.Output['outputs.NetworkpolicyGuest']:
        """
        Guest. The structure of `guest` block is documented below.
        """
        return pulumi.get(self, "guest")

    @property
    @pulumi.getter(name="guestVoiceSignaling")
    def guest_voice_signaling(self) -> pulumi.Output['outputs.NetworkpolicyGuestVoiceSignaling']:
        """
        Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
        """
        return pulumi.get(self, "guest_voice_signaling")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        LLDP network policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def softphone(self) -> pulumi.Output['outputs.NetworkpolicySoftphone']:
        """
        Softphone. The structure of `softphone` block is documented below.
        """
        return pulumi.get(self, "softphone")

    @property
    @pulumi.getter(name="streamingVideo")
    def streaming_video(self) -> pulumi.Output['outputs.NetworkpolicyStreamingVideo']:
        """
        Streaming Video. The structure of `streaming_video` block is documented below.
        """
        return pulumi.get(self, "streaming_video")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="videoConferencing")
    def video_conferencing(self) -> pulumi.Output['outputs.NetworkpolicyVideoConferencing']:
        """
        Video Conferencing. The structure of `video_conferencing` block is documented below.
        """
        return pulumi.get(self, "video_conferencing")

    @property
    @pulumi.getter(name="videoSignaling")
    def video_signaling(self) -> pulumi.Output['outputs.NetworkpolicyVideoSignaling']:
        """
        Video Signaling. The structure of `video_signaling` block is documented below.
        """
        return pulumi.get(self, "video_signaling")

    @property
    @pulumi.getter
    def voice(self) -> pulumi.Output['outputs.NetworkpolicyVoice']:
        """
        Voice. The structure of `voice` block is documented below.
        """
        return pulumi.get(self, "voice")

    @property
    @pulumi.getter(name="voiceSignaling")
    def voice_signaling(self) -> pulumi.Output['outputs.NetworkpolicyVoiceSignaling']:
        """
        Voice signaling. The structure of `voice_signaling` block is documented below.
        """
        return pulumi.get(self, "voice_signaling")

