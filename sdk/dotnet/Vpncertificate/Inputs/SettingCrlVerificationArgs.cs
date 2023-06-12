// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpncertificate.Inputs
{

    public sealed class SettingCrlVerificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CRL verification option when CRL of any certificate in chain is absent (default = ignore). Valid values: `ignore`, `revoke`.
        /// </summary>
        [Input("chainCrlAbsence")]
        public Input<string>? ChainCrlAbsence { get; set; }

        /// <summary>
        /// CRL verification option when CRL is expired (default = ignore). Valid values: `ignore`, `revoke`.
        /// </summary>
        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        /// <summary>
        /// CRL verification option when leaf CRL is absent (default = ignore). Valid values: `ignore`, `revoke`.
        /// </summary>
        [Input("leafCrlAbsence")]
        public Input<string>? LeafCrlAbsence { get; set; }

        public SettingCrlVerificationArgs()
        {
        }
        public static new SettingCrlVerificationArgs Empty => new SettingCrlVerificationArgs();
    }
}
