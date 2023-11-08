import React, { type MouseEvent, useMemo, useState, useCallback, useLayoutEffect, Fragment, type FC } from 'react'

import { mdiInformationOutline } from '@mdi/js'
import classNames from 'classnames'

import { isSafari } from '@sourcegraph/common'
import { shortcutDisplayName } from '@sourcegraph/shared/src/keyboardShortcuts'
import { Icon, useWindowSize } from '@sourcegraph/wildcard'

import { decorateQuery } from '../../util/query'

import { RenderAs, type Action, type Group, type Option } from './suggestionsExtension'
import { getSpans } from './utils'

import styles from './Suggestions.module.scss'

type Renderable = React.ReactElement | string | null

function getActionName(action: Action): string {
    switch (action.type) {
        case 'completion': {
            return action.name ?? 'Add'
        }
        case 'goto': {
            return action.name ?? 'Go to'
        }
        case 'command': {
            return action.name ?? 'Run'
        }
    }
}

interface SuggestionsProps {
    id: string
    results: Group[]
    activeRowIndex: number
    open?: boolean
    onSelect(option: Option, action?: Action): void
}

export const Suggestions: FC<SuggestionsProps> = ({ id, results, activeRowIndex, onSelect, open = false }) => {
    const [container, setContainer] = useState<HTMLDivElement | null>(null)

    // Handles mouse clicks on suggestions. The corresponding option is determined by the extracting group and option
    // indicies from the element ID.
    const handleSelection = useCallback(
        (event: MouseEvent) => {
            const target = event.target as HTMLElement
            const match = target.closest('li[role="row"]')?.id.match(/\d+x\d+/)
            if (match) {
                // Extracts the group and row index from the elements ID to pass
                // the right option value to the callback.
                const [groupIndex, optionIndex] = match[0].split('x')
                const option = results[+groupIndex].options[+optionIndex]
                // Determine which action was selected.
                onSelect(
                    option,
                    target.closest<HTMLElement>('[data-action]')?.dataset?.action === 'secondary'
                        ? option.alternativeAction
                        : option.action
                )
            }
        },
        [onSelect, results]
    )

    const { height: windowHeight } = useWindowSize()
    const maxHeight = useMemo(
        // This is using an arbitrary 20px "margin" between the suggestions box
        // and the window border
        () => (container ? `${windowHeight - container.getBoundingClientRect().top - 20}px` : 'auto'),
        // Recompute height when suggestions change
        // eslint-disable-next-line react-hooks/exhaustive-deps
        [container, windowHeight, results]
    )
    const flattenedRows = useMemo(() => results.flatMap(group => group.options), [results])
    const focusedItem = flattenedRows[activeRowIndex]
    const show = open && results.length > 0

    useLayoutEffect(() => {
        if (container) {
            // Options are not supported in Safari according to
            // https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView#browser_compatibility
            container.querySelector('[aria-selected="true"]')?.scrollIntoView(isSafari() ? false : { block: 'nearest' })
        }
    }, [container, focusedItem])

    if (!show) {
        return null
    }

    return (
        <div
            ref={setContainer}
            id={id}
            className={styles.container}
            // eslint-disable-next-line react/forbid-dom-props
            style={{ maxHeight }}
        >
            <div className={styles.suggestions} role="grid" onMouseDown={handleSelection} tabIndex={-1}>
                {results.map((group, groupIndex) =>
                    group.options.length > 0 ? (
                        <ul role="rowgroup" key={group.title} aria-labelledby={`${id}-${groupIndex}-label`}>
                            <li id={`${id}-${groupIndex}-label`} role="presentation">
                                {group.title}
                            </li>
                            {group.options.map((option, rowIndex) => (
                                <li
                                    role="row"
                                    key={rowIndex}
                                    id={`${id}-${groupIndex}x${rowIndex}`}
                                    aria-selected={focusedItem === option}
                                >
                                    {option.icon && (
                                        <div className="pr-1 align-self-start">
                                            <Icon className={styles.icon} svgPath={option.icon} aria-hidden="true" />
                                        </div>
                                    )}
                                    <div className={styles.innerRow}>
                                        <div className="d-flex flex-wrap">
                                            <div
                                                role="gridcell"
                                                className={classNames(styles.label, 'test-option-label')}
                                            >
                                                {option.render === undefined && (
                                                    <EmphasizedLabel label={option.label} matches={option.matches} />
                                                )}
                                                {option.render === RenderAs.FILTER && <FilterOption option={option} />}
                                                {option.render === RenderAs.QUERY && <QueryOption option={option} />}
                                            </div>
                                            {option.description && (
                                                <div role="gridcell" className={styles.description}>
                                                    {option.description}
                                                </div>
                                            )}
                                        </div>
                                        <div className={styles.note}>
                                            <div role="gridcell" data-action="primary">
                                                {getActionName(option.action)}
                                            </div>
                                            {option.alternativeAction && (
                                                <div role="gridcell" data-action="secondary">
                                                    {getActionName(option.alternativeAction)}
                                                </div>
                                            )}
                                        </div>
                                    </div>
                                </li>
                            ))}
                        </ul>
                    ) : null
                )}
            </div>
            {focusedItem && <Footer option={focusedItem} />}
        </div>
    )
}

const Footer: FC<{ option: Option }> = ({ option }) => (
    <div className={classNames(styles.footer, 'd-flex align-items-center justify-content-between')}>
        <span>
            <ActionInfo action={option.action} shortcut="Enter" />{' '}
            {option.alternativeAction && <ActionInfo action={option.alternativeAction} shortcut="Mod+Enter" />}
        </span>
        <Icon className={styles.icon} svgPath={mdiInformationOutline} aria-hidden="true" />
    </div>
)

const ActionInfo: FC<{ action: Action; shortcut: string }> = ({ action, shortcut }) => {
    let info: Renderable = action.info ?? null
    switch (action.type) {
        case 'completion': {
            info = (
                <>
                    <strong>add</strong> to your query
                </>
            )
            break
        }
        case 'goto': {
            info = (
                <>
                    <strong>go to</strong> the suggestion
                </>
            )
            break
        }
        case 'command': {
            info = (
                <>
                    <strong>execute</strong> the command
                </>
            )
            break
        }
    }

    return (
        <>
            Press <kbd>{shortcutDisplayName(shortcut)}</kbd> to {info}.
        </>
    )
}

const FilterOption: FC<{ option: Option }> = ({ option }) => {
    const label = option.label
    const separatorIndex = label.indexOf(':')
    const field = separatorIndex > -1 ? label.slice(0, separatorIndex) : label
    const value = separatorIndex > -1 ? label.slice(separatorIndex + 1) : ''

    return (
        <span className={styles.filterOption}>
            <span className="search-filter-keyword">
                {option.matches ? <EmphasizedLabel label={field} matches={option.matches} /> : field}
                <span className={styles.separator}>:</span>
            </span>
            {value && <EmphasizedLabel label={value} matches={option.matches} offset={field.length + 1} />}
        </span>
    )
}

/**
 * Adaption of {@link SyntaxHighlightedSearchQuery} to support emphasizing matching substrings.
 */
const QueryOption: FC<{ option: Option }> = ({ option: { label, matches } }) => {
    const tokens = useMemo(() => {
        const decorations = decorateQuery(label)

        return decorations
            ? decorations.map(({ value, key, className, token }) => (
                  <span className={className} key={key}>
                      <EmphasizedLabel label={value} matches={matches} offset={token.range.start} />
                  </span>
              ))
            : [<Fragment key="0">{label}</Fragment>]
    }, [label, matches])

    return <span className="text-monospace search-query-link">{tokens}</span>
}

/**
 * Emphasizes matching substring in the provided label.
 */
const EmphasizedLabel: FC<{ label: string; matches?: Set<number>; offset?: number }> = ({
    label,
    matches,
    offset = 0,
}) => {
    if (!matches) {
        return <span>{label}</span>
    }
    const spans = getSpans(matches, label.length, offset)
    return (
        <span>
            {spans.map(([start, end, match]) => {
                const value = label.slice(start, end + 1)
                return match ? (
                    <span key={offset + start} className={styles.match}>
                        {value}
                    </span>
                ) : (
                    value
                )
            })}
        </span>
    )
}
