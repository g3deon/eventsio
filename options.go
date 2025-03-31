package eventsio

type PublishOption func(*PublishOptions)

type PublishOptions struct {
	isRetry bool
}

func WithRetry() PublishOption {
	return func(o *PublishOptions) {
		o.isRetry = true
	}
}

func (o *PublishOptions) IsRetry() bool {
	return o.isRetry
}
