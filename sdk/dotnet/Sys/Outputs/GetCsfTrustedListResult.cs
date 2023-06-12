// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GetCsfTrustedListResult
    {
        /// <summary>
        /// Security fabric authorization action.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Authorization type.
        /// </summary>
        public readonly string AuthorizationType;
        /// <summary>
        /// Certificate.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Trust authorizations by this node's administrator.
        /// </summary>
        public readonly string DownstreamAuthorization;
        /// <summary>
        /// HA members.
        /// </summary>
        public readonly string HaMembers;
        /// <summary>
        /// Device name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Serial.
        /// </summary>
        public readonly string Serial;

        [OutputConstructor]
        private GetCsfTrustedListResult(
            string action,

            string authorizationType,

            string certificate,

            string downstreamAuthorization,

            string haMembers,

            string name,

            string serial)
        {
            Action = action;
            AuthorizationType = authorizationType;
            Certificate = certificate;
            DownstreamAuthorization = downstreamAuthorization;
            HaMembers = haMembers;
            Name = name;
            Serial = serial;
        }
    }
}
