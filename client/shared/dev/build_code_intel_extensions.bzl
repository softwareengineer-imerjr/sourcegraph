"Schema generation rule"

load("@aspect_rules_js//js:defs.bzl", "js_run_binary")

def build_code_intel_extensions(name, out, revision):
    """ Download code-intel extension bundles from GitHub.

    Args:
        name: target name
        out: output revisions folder
        revision: revision
    """
    js_run_binary(
        name = name,
        chdir = native.package_name(),
        out_dirs = [out],
        log_level = "info",
        silent_on_success = False,
        args = [
            revision,
            out,
            # "$(location @curl_nix//:bin/curl)",
        ],
        srcs = [
            "@curl_nix//:bin",
            "@bash_nix//:bin",
        ],
        tags = [
            # We download static assets from GitHub.
            "requires-network",
        ],
        env = {
            "PATH": "$(locations @bash_nix//:bin):$(location @curl_nix//:bin)",
        },
        copy_srcs_to_bin = False,
        tool = "//client/shared/dev:build_code_intel_extensions",
    )
