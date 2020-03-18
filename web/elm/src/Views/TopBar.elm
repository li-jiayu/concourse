module Views.TopBar exposing
    ( breadcrumbs
    , concourseLogo
    )

import Assets
import Concourse
import Html exposing (Html)
import Html.Attributes
    exposing
        ( class
        , href
        , id
        )
import Message.Message exposing (DomID(..), Message(..))
import Routes
import Url
import Views.Styles as Styles


concourseLogo : Html Message
concourseLogo =
    Html.a (href "/" :: Styles.concourseLogo) []


breadcrumbs : Routes.Route -> Html Message
breadcrumbs route =
    Html.div
        (id "breadcrumbs" :: Styles.breadcrumbContainer)
    <|
        case route of
            Routes.Pipeline { id } ->
                [ pipelineBreadcumb
                    { teamName = id.teamName
                    , pipelineName = id.pipelineName
                    }
                ]

            Routes.Build { id } ->
                [ pipelineBreadcumb
                    { teamName = id.teamName
                    , pipelineName = id.pipelineName
                    }
                , breadcrumbSeparator
                , jobBreadcrumb id.jobName
                ]

            Routes.Resource { id } ->
                [ pipelineBreadcumb
                    { teamName = id.teamName
                    , pipelineName = id.pipelineName
                    }
                , breadcrumbSeparator
                , resourceBreadcrumb id.resourceName
                ]

            Routes.Job { id } ->
                [ pipelineBreadcumb
                    { teamName = id.teamName
                    , pipelineName = id.pipelineName
                    }
                , breadcrumbSeparator
                , jobBreadcrumb id.jobName
                ]

            _ ->
                []


breadcrumbComponent : Assets.ComponentType -> String -> List (Html Message)
breadcrumbComponent componentType name =
    [ Html.div
        (Styles.breadcrumbComponent componentType)
        []
    , Html.text <| decodeName name
    ]


breadcrumbSeparator : Html Message
breadcrumbSeparator =
    Html.li
        (class "breadcrumb-separator" :: Styles.breadcrumbItem False)
        [ Html.text "/" ]


pipelineBreadcumb : Concourse.PipelineIdentifier -> Html Message
pipelineBreadcumb pipelineId =
    Html.a
        ([ id "breadcrumb-pipeline"
         , href <|
            Routes.toString <|
                Routes.Pipeline { id = pipelineId, groups = [] }
         ]
            ++ Styles.breadcrumbItem True
        )
        (breadcrumbComponent Assets.PipelineComponent pipelineId.pipelineName)


jobBreadcrumb : String -> Html Message
jobBreadcrumb jobName =
    Html.li
        (id "breadcrumb-job" :: Styles.breadcrumbItem False)
        (breadcrumbComponent Assets.JobComponent jobName)


resourceBreadcrumb : String -> Html Message
resourceBreadcrumb resourceName =
    Html.li
        (id "breadcrumb-resource" :: Styles.breadcrumbItem False)
        (breadcrumbComponent Assets.ResourceComponent resourceName)


decodeName : String -> String
decodeName name =
    Maybe.withDefault name (Url.percentDecode name)
