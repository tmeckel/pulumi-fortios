// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetIsisIsisNetResult
    {
        /// <summary>
        /// Prefix entry ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// IS-IS net xx.xxxx. ... .xxxx.xx.
        /// </summary>
        public readonly string Net;

        [OutputConstructor]
        private GetIsisIsisNetResult(
            int id,

            string net)
        {
            Id = id;
            Net = net;
        }
    }
}
