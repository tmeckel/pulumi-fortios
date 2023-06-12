// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class AddrgrpTaggingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Tagging entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.AddrgrpTaggingTagGetArgs>? _tags;

        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpTaggingTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AddrgrpTaggingTagGetArgs>());
            set => _tags = value;
        }

        public AddrgrpTaggingGetArgs()
        {
        }
        public static new AddrgrpTaggingGetArgs Empty => new AddrgrpTaggingGetArgs();
    }
}
