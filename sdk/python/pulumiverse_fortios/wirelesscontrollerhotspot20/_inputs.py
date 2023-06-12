# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'Anqp3gppcellularMccMncListArgs',
    'AnqpnairealmNaiListArgs',
    'AnqpnairealmNaiListEapMethodArgs',
    'AnqpnairealmNaiListEapMethodAuthParamArgs',
    'AnqproamingconsortiumOiListArgs',
    'AnqpvenuenameValueListArgs',
    'AnqpvenueurlValueListArgs',
    'H2qpadviceofchargeAocListArgs',
    'H2qpadviceofchargeAocListPlanInfoArgs',
    'H2qpoperatornameValueListArgs',
    'H2qposuproviderFriendlyNameArgs',
    'H2qposuproviderServiceDescriptionArgs',
    'H2qposuprovidernaiNaiListArgs',
    'HsprofileOsuProviderArgs',
    'IconIconListArgs',
    'QosmapDscpExceptArgs',
    'QosmapDscpRangeArgs',
]

@pulumi.input_type
class Anqp3gppcellularMccMncListArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[int]] = None,
                 mcc: Optional[pulumi.Input[str]] = None,
                 mnc: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] mcc: Mobile country code.
        :param pulumi.Input[str] mnc: Mobile network code.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if mcc is not None:
            pulumi.set(__self__, "mcc", mcc)
        if mnc is not None:
            pulumi.set(__self__, "mnc", mnc)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def mcc(self) -> Optional[pulumi.Input[str]]:
        """
        Mobile country code.
        """
        return pulumi.get(self, "mcc")

    @mcc.setter
    def mcc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mcc", value)

    @property
    @pulumi.getter
    def mnc(self) -> Optional[pulumi.Input[str]]:
        """
        Mobile network code.
        """
        return pulumi.get(self, "mnc")

    @mnc.setter
    def mnc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mnc", value)


@pulumi.input_type
class AnqpnairealmNaiListArgs:
    def __init__(__self__, *,
                 eap_methods: Optional[pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodArgs']]]] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 nai_realm: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodArgs']]] eap_methods: EAP Methods. The structure of `eap_method` block is documented below.
        :param pulumi.Input[str] encoding: Enable/disable format in accordance with IETF RFC 4282. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] nai_realm: Configure NAI realms (delimited by a semi-colon character).
        :param pulumi.Input[str] name: NAI realm name.
        """
        if eap_methods is not None:
            pulumi.set(__self__, "eap_methods", eap_methods)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if nai_realm is not None:
            pulumi.set(__self__, "nai_realm", nai_realm)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="eapMethods")
    def eap_methods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodArgs']]]]:
        """
        EAP Methods. The structure of `eap_method` block is documented below.
        """
        return pulumi.get(self, "eap_methods")

    @eap_methods.setter
    def eap_methods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodArgs']]]]):
        pulumi.set(self, "eap_methods", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable format in accordance with IETF RFC 4282. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="naiRealm")
    def nai_realm(self) -> Optional[pulumi.Input[str]]:
        """
        Configure NAI realms (delimited by a semi-colon character).
        """
        return pulumi.get(self, "nai_realm")

    @nai_realm.setter
    def nai_realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nai_realm", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        NAI realm name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class AnqpnairealmNaiListEapMethodArgs:
    def __init__(__self__, *,
                 auth_params: Optional[pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodAuthParamArgs']]]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 method: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodAuthParamArgs']]] auth_params: EAP auth param. The structure of `auth_param` block is documented below.
        :param pulumi.Input[int] index: EAP method index.
        :param pulumi.Input[str] method: EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.
        """
        if auth_params is not None:
            pulumi.set(__self__, "auth_params", auth_params)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if method is not None:
            pulumi.set(__self__, "method", method)

    @property
    @pulumi.getter(name="authParams")
    def auth_params(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodAuthParamArgs']]]]:
        """
        EAP auth param. The structure of `auth_param` block is documented below.
        """
        return pulumi.get(self, "auth_params")

    @auth_params.setter
    def auth_params(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AnqpnairealmNaiListEapMethodAuthParamArgs']]]]):
        pulumi.set(self, "auth_params", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        EAP method index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)


@pulumi.input_type
class AnqpnairealmNaiListEapMethodAuthParamArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 val: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.
        :param pulumi.Input[int] index: Param index.
        :param pulumi.Input[str] val: Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if val is not None:
            pulumi.set(__self__, "val", val)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        Param index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def val(self) -> Optional[pulumi.Input[str]]:
        """
        Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.
        """
        return pulumi.get(self, "val")

    @val.setter
    def val(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "val", value)


@pulumi.input_type
class AnqproamingconsortiumOiListArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 oi: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] index: OI index.
        :param pulumi.Input[str] oi: Organization identifier.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if oi is not None:
            pulumi.set(__self__, "oi", oi)

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
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        OI index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def oi(self) -> Optional[pulumi.Input[str]]:
        """
        Organization identifier.
        """
        return pulumi.get(self, "oi")

    @oi.setter
    def oi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oi", value)


@pulumi.input_type
class AnqpvenuenameValueListArgs:
    def __init__(__self__, *,
                 index: Optional[pulumi.Input[int]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] index: Value index.
        :param pulumi.Input[str] lang: Language code.
        :param pulumi.Input[str] value: Venue name value.
        """
        if index is not None:
            pulumi.set(__self__, "index", index)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        Value index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        Language code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Venue name value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class AnqpvenueurlValueListArgs:
    def __init__(__self__, *,
                 index: Optional[pulumi.Input[int]] = None,
                 number: Optional[pulumi.Input[int]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] index: URL index.
        :param pulumi.Input[int] number: Venue number.
        :param pulumi.Input[str] value: Venue URL value.
        """
        if index is not None:
            pulumi.set(__self__, "index", index)
        if number is not None:
            pulumi.set(__self__, "number", number)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        URL index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def number(self) -> Optional[pulumi.Input[int]]:
        """
        Venue number.
        """
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Venue URL value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class H2qpadviceofchargeAocListArgs:
    def __init__(__self__, *,
                 nai_realm: Optional[pulumi.Input[str]] = None,
                 nai_realm_encoding: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan_infos: Optional[pulumi.Input[Sequence[pulumi.Input['H2qpadviceofchargeAocListPlanInfoArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] nai_realm: NAI realm list name.
        :param pulumi.Input[str] nai_realm_encoding: NAI realm encoding.
        :param pulumi.Input[str] name: Advice of charge ID.
        :param pulumi.Input[Sequence[pulumi.Input['H2qpadviceofchargeAocListPlanInfoArgs']]] plan_infos: Plan info. The structure of `plan_info` block is documented below.
        :param pulumi.Input[str] type: Usage charge type. Valid values: `time-based`, `volume-based`, `time-and-volume-based`, `unlimited`.
        """
        if nai_realm is not None:
            pulumi.set(__self__, "nai_realm", nai_realm)
        if nai_realm_encoding is not None:
            pulumi.set(__self__, "nai_realm_encoding", nai_realm_encoding)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plan_infos is not None:
            pulumi.set(__self__, "plan_infos", plan_infos)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="naiRealm")
    def nai_realm(self) -> Optional[pulumi.Input[str]]:
        """
        NAI realm list name.
        """
        return pulumi.get(self, "nai_realm")

    @nai_realm.setter
    def nai_realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nai_realm", value)

    @property
    @pulumi.getter(name="naiRealmEncoding")
    def nai_realm_encoding(self) -> Optional[pulumi.Input[str]]:
        """
        NAI realm encoding.
        """
        return pulumi.get(self, "nai_realm_encoding")

    @nai_realm_encoding.setter
    def nai_realm_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nai_realm_encoding", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Advice of charge ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="planInfos")
    def plan_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['H2qpadviceofchargeAocListPlanInfoArgs']]]]:
        """
        Plan info. The structure of `plan_info` block is documented below.
        """
        return pulumi.get(self, "plan_infos")

    @plan_infos.setter
    def plan_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['H2qpadviceofchargeAocListPlanInfoArgs']]]]):
        pulumi.set(self, "plan_infos", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Usage charge type. Valid values: `time-based`, `volume-based`, `time-and-volume-based`, `unlimited`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class H2qpadviceofchargeAocListPlanInfoArgs:
    def __init__(__self__, *,
                 currency: Optional[pulumi.Input[str]] = None,
                 info_file: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] currency: Currency code.
        :param pulumi.Input[str] info_file: Info file.
        :param pulumi.Input[str] lang: Languague code.
        :param pulumi.Input[str] name: Plan name.
        """
        if currency is not None:
            pulumi.set(__self__, "currency", currency)
        if info_file is not None:
            pulumi.set(__self__, "info_file", info_file)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def currency(self) -> Optional[pulumi.Input[str]]:
        """
        Currency code.
        """
        return pulumi.get(self, "currency")

    @currency.setter
    def currency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "currency", value)

    @property
    @pulumi.getter(name="infoFile")
    def info_file(self) -> Optional[pulumi.Input[str]]:
        """
        Info file.
        """
        return pulumi.get(self, "info_file")

    @info_file.setter
    def info_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "info_file", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        Languague code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Plan name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class H2qpoperatornameValueListArgs:
    def __init__(__self__, *,
                 index: Optional[pulumi.Input[int]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] index: Value index.
        :param pulumi.Input[str] lang: Language code.
        :param pulumi.Input[str] value: Friendly name value.
        """
        if index is not None:
            pulumi.set(__self__, "index", index)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        Value index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        Language code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Friendly name value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class H2qposuproviderFriendlyNameArgs:
    def __init__(__self__, *,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 lang: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] friendly_name: OSU provider friendly name.
        :param pulumi.Input[int] index: OSU provider friendly name index.
        :param pulumi.Input[str] lang: Language code.
        """
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        """
        OSU provider friendly name.
        """
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        OSU provider friendly name index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        Language code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)


@pulumi.input_type
class H2qposuproviderServiceDescriptionArgs:
    def __init__(__self__, *,
                 lang: Optional[pulumi.Input[str]] = None,
                 service_description: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] lang: Language code.
        :param pulumi.Input[str] service_description: Service description.
        :param pulumi.Input[int] service_id: OSU service ID.
        """
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if service_description is not None:
            pulumi.set(__self__, "service_description", service_description)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        Language code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="serviceDescription")
    def service_description(self) -> Optional[pulumi.Input[str]]:
        """
        Service description.
        """
        return pulumi.get(self, "service_description")

    @service_description.setter
    def service_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_description", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[int]]:
        """
        OSU service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "service_id", value)


@pulumi.input_type
class H2qposuprovidernaiNaiListArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 osu_nai: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: OSU NAI ID.
        :param pulumi.Input[str] osu_nai: OSU NAI.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if osu_nai is not None:
            pulumi.set(__self__, "osu_nai", osu_nai)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        OSU NAI ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osuNai")
    def osu_nai(self) -> Optional[pulumi.Input[str]]:
        """
        OSU NAI.
        """
        return pulumi.get(self, "osu_nai")

    @osu_nai.setter
    def osu_nai(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "osu_nai", value)


@pulumi.input_type
class HsprofileOsuProviderArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: OSU provider name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        OSU provider name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class IconIconListArgs:
    def __init__(__self__, *,
                 file: Optional[pulumi.Input[str]] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] file: Icon file.
        :param pulumi.Input[int] height: Icon height.
        :param pulumi.Input[str] lang: Language code.
        :param pulumi.Input[str] name: Icon name.
        :param pulumi.Input[str] type: Icon type. Valid values: `bmp`, `gif`, `jpeg`, `png`, `tiff`.
        :param pulumi.Input[int] width: Icon width.
        """
        if file is not None:
            pulumi.set(__self__, "file", file)
        if height is not None:
            pulumi.set(__self__, "height", height)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if width is not None:
            pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        Icon file.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def height(self) -> Optional[pulumi.Input[int]]:
        """
        Icon height.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        Language code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Icon name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Icon type. Valid values: `bmp`, `gif`, `jpeg`, `png`, `tiff`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        Icon width.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)


@pulumi.input_type
class QosmapDscpExceptArgs:
    def __init__(__self__, *,
                 dscp: Optional[pulumi.Input[int]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 up: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] dscp: DSCP value.
        :param pulumi.Input[int] index: DSCP exception index.
        :param pulumi.Input[int] up: User priority.
        """
        if dscp is not None:
            pulumi.set(__self__, "dscp", dscp)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if up is not None:
            pulumi.set(__self__, "up", up)

    @property
    @pulumi.getter
    def dscp(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP value.
        """
        return pulumi.get(self, "dscp")

    @dscp.setter
    def dscp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dscp", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP exception index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def up(self) -> Optional[pulumi.Input[int]]:
        """
        User priority.
        """
        return pulumi.get(self, "up")

    @up.setter
    def up(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "up", value)


@pulumi.input_type
class QosmapDscpRangeArgs:
    def __init__(__self__, *,
                 high: Optional[pulumi.Input[int]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 low: Optional[pulumi.Input[int]] = None,
                 up: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] high: DSCP high value.
        :param pulumi.Input[int] index: DSCP range index.
        :param pulumi.Input[int] low: DSCP low value.
        :param pulumi.Input[int] up: User priority.
        """
        if high is not None:
            pulumi.set(__self__, "high", high)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if low is not None:
            pulumi.set(__self__, "low", low)
        if up is not None:
            pulumi.set(__self__, "up", up)

    @property
    @pulumi.getter
    def high(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP high value.
        """
        return pulumi.get(self, "high")

    @high.setter
    def high(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "high", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP range index.
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def low(self) -> Optional[pulumi.Input[int]]:
        """
        DSCP low value.
        """
        return pulumi.get(self, "low")

    @low.setter
    def low(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "low", value)

    @property
    @pulumi.getter
    def up(self) -> Optional[pulumi.Input[int]]:
        """
        User priority.
        """
        return pulumi.get(self, "up")

    @up.setter
    def up(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "up", value)

