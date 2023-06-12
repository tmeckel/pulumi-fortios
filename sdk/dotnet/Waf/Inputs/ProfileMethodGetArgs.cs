// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Inputs
{

    public sealed class ProfileMethodGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Methods. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`, `others`.
        /// </summary>
        [Input("defaultAllowedMethods")]
        public Input<string>? DefaultAllowedMethods { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        [Input("methodPolicies")]
        private InputList<Inputs.ProfileMethodMethodPolicyGetArgs>? _methodPolicies;

        /// <summary>
        /// HTTP method policy. The structure of `method_policy` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileMethodMethodPolicyGetArgs> MethodPolicies
        {
            get => _methodPolicies ?? (_methodPolicies = new InputList<Inputs.ProfileMethodMethodPolicyGetArgs>());
            set => _methodPolicies = value;
        }

        /// <summary>
        /// Severity. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProfileMethodGetArgs()
        {
        }
        public static new ProfileMethodGetArgs Empty => new ProfileMethodGetArgs();
    }
}