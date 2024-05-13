<script lang="ts">
    import { mdiFolder } from '@mdi/js'
    import { onMount } from 'svelte'

    import Avatar from '$lib/Avatar.svelte'
    import { getGraphQLClient } from '$lib/graphql'
    import Icon from '$lib/Icon.svelte'
    import { displayRepoName } from '$lib/shared'
    import Timestamp from '$lib/Timestamp.svelte'
    import { formatBytes } from '$lib/utils'

    import { FileOrDirPopoverQuery } from '../../../routes/[...repo=reporev]/(validrev)/(code)/layout.gql'
    import type { FileIcon_GitBlob } from '../FileIcon.gql'
    import FileIcon from '../FileIcon.svelte'

    import { type DirPopoverFragment, type FilePopoverFragment } from './FilePopover.gql'
    import NodeLine from './NodeLine.svelte'

    export let repoName: string
    export let revspec: string
    export let filePath: string
    export let entry: FileIcon_GitBlob
    let fileFrag: FilePopoverFragment | null = null
    let dirFrag: DirPopoverFragment | null = null

    const fetchPopoverData = async () => {
        const client = getGraphQLClient()
        const result = await client.query(FileOrDirPopoverQuery, {
            repoName,
            revspec,
            filePath,
        })

        if (result.error || result.data === undefined) {
            console.error('could not fetch file or dir popover', result.error)
            throw new Error('could not fetch file or dir popover', result.error)
        }

        if (result?.data?.repository?.commit?.path?.__typename === 'GitBlob') {
            fileFrag = result?.data?.repository?.commit?.path
            return fileFrag
        } else if (result?.data?.repository?.commit?.path?.__typename === 'GitTree') {
            dirFrag = result?.data?.repository?.commit?.path
            return dirFrag
        }

        return null
    }

    const abbreviatedFilePath = (filePath: string): string => {
        let parts = filePath.split('/')
        if (parts.length <= 3) return parts.join(' / ')

        return `${parts[0]} / ... / ${parts[parts.length - 1]}`
    }

    const getFileName = (filePath: string): string => {
        let parts = filePath.split('/')
        return parts[parts.length - 1]
    }

    const CENTER_DOT = '\u00B7' // interpunct

    onMount(async () => {
        fetchPopoverData()
    })

    $: abbreviatedPath = abbreviatedFilePath(filePath)
    $: repo = displayRepoName(repoName).replace('/', ' / ')
    $: repoAndPath = `${repo} ${CENTER_DOT} ${abbreviatedPath}`
    $: fileOrDirName = getFileName(filePath)
    $: fileCommit = fileFrag?.blame[0].commit
    $: languages = fileFrag?.languages[0]
</script>

{#if fileFrag || dirFrag}
    <div class="root">
        <div class="desc">
            <div class="repo-and-path">
                <small>{repoAndPath}</small>
            </div>

            <div class="lang-and-file">
                {#if dirFrag}
                    <Icon svgPath={mdiFolder} --icon-fill-color="var(--primary)" --icon-size="1.5rem" />
                {:else if fileFrag}
                    <FileIcon file={entry} inline={false} size="1.5rem" />
                {/if}
                <div class="file">
                    <div>{fileOrDirName}</div>
                    {#if fileFrag && !dirFrag}
                        <small
                            >{languages ? fileFrag.languages[0] : ''}
                            {languages ? CENTER_DOT : ''}
                            {fileFrag.totalLines} Lines {CENTER_DOT} Bytes {formatBytes(fileFrag.byteSize)}</small
                        >
                    {:else if dirFrag && !fileFrag}
                        <small
                            >Subdirectories {dirFrag.directories.length}
                            {CENTER_DOT} Files {dirFrag.files.length}</small
                        >
                    {/if}
                </div>
            </div>
        </div>

        <div class="last-commit">
            <small class="title">Last Changed @</small>
            <div class="commit">
                <NodeLine />
                {#if fileFrag && !dirFrag}
                    <div>
                        <a href={fileCommit?.canonicalURL} target="_blank">
                            {fileCommit?.abbreviatedOID}
                        </a>
                        <div class="msg">{fileCommit?.subject}</div>
                        <div class="author">
                            <Avatar avatar={fileFrag.blame[0].author.person} --avatar-size="1.0rem" />
                            <small class="name"
                                >{fileFrag?.blame[0]?.author?.person?.displayName}
                                {CENTER_DOT}
                                <Timestamp date={fileFrag?.blame[0]?.author.date} /></small
                            >
                        </div>
                    </div>
                {:else if dirFrag && !fileFrag}
                    <div>
                        <a href={dirFrag?.commit?.canonicalURL} target="_blank">
                            {dirFrag?.commit?.abbreviatedOID}
                        </a>
                        <div class="msg">{dirFrag?.commit.subject}</div>
                        <div class="author">
                            <Avatar avatar={dirFrag?.commit.author.person} --avatar-size="1.0rem" />
                            <small class="name"
                                >{dirFrag?.commit?.author?.person?.displayName}
                                {CENTER_DOT}
                                <Timestamp date={dirFrag?.commit?.author?.date} /></small
                            >
                        </div>
                    </div>
                {/if}
            </div>
        </div>
    </div>
{/if}

<style lang="scss">
    .root {
        width: 400px;
        background: var(--body-bg);
        border: 1px solid var(--border-color);
        border-radius: 8px;

        .desc {
            display: flex;
            flex-flow: column nowrap;
            align-items: center;
            justify-content: center;

            .repo-and-path {
                align-items: center;
                border-bottom: 1px solid var(--border-color);
                display: flex;
                flex-flow: row nowrap;
                gap: 0.25rem;
                justify-content: flex-start;
                padding: 0.5rem 1rem;
                width: 100%;
                font-family: var(--monospace-font-family);
                white-space: nowrap;
                color: var(--text-muted);

                small {
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }
            }

            .lang-and-file {
                width: 100%;
                display: flex;
                flex-flow: row nowrap;
                align-items: center;
                justify-content: flex-start;
                padding: 0.5rem 1rem;
                gap: 0.25rem 0.75rem;

                .file {
                    display: flex;
                    flex-flow: column nowrap;
                    align-items: flex-start;
                    justify-content: flex-start;
                    font-family: var(--monospace-font-family);
                    gap: 0.25rem;

                    div {
                        color: var(--text-body);
                    }

                    small {
                        color: var(--text-muted);
                    }
                }
            }
        }

        .last-commit {
            display: flex;
            flex-flow: column nowrap;
            align-items: flex-start;
            justify-content: center;
            gap: 0.5rem 0.5rem 0rem;

            .title {
                padding: 0.5rem 1rem;
                color: var(--text-body);
                background-color: var(--secondary-4);
                width: 100%;
                border-bottom: 1px solid var(--border-color);
            }

            .commit {
                padding-left: 1.5rem;
                display: flex;
                flex-flow: row nowrap;
                align-items: center;
                justify-content: flex-start;
                width: 100%;
                height: 90px;
                gap: 0.5rem 1.25rem;

                div {
                    display: flex;
                    flex-flow: column nowrap;
                    align-items: flex-start;
                    justify-content: center;
                    gap: 0.25rem;
                    width: 325px;

                    a {
                        font-family: var(--monospace-font-family);
                        background-color: var(--color-bg-2);
                        padding: 0.15rem 0.25rem;
                        border-radius: 3px;
                    }

                    .msg {
                        color: var(--text-body);
                        text-overflow: ellipsis ellipsis;
                        overflow: hidden;
                        white-space: nowrap;
                    }
                    .author {
                        display: flex;
                        flex-flow: row nowrap;
                        justify-content: flex-start;
                        align-items: center;
                        gap: 0.25rem 0.5rem;
                        color: var(--text-muted);
                        .name {
                            margin-right: 0.5rem;
                        }
                    }
                }
            }
        }
    }
</style>
